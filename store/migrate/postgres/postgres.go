package postgres

import (
	"database/sql"

	"github.com/apex/log"
	"github.com/jmoiron/sqlx"
)

var migrations = []struct {
	name string
	stmt string
}{
	{name: "001-create-buyer-table", stmt: createBuyerTable},
	{name: "002-create-receiver-address-table", stmt: createReceiverAddressTable},
	{name: "003-create-shipping-table", stmt: createShippingTable},
	{name: "004-create-orders-table", stmt: createOrdersTable},
	{name: "005-create-order-item-table", stmt: createOrderItemTable},
	{name: "006-create-payment-table", stmt: createPaymentTable},
}

// Migrate performs the database migration. If the migration fails
// and error is returned.
func Migrate(db *sqlx.DB) error {
	if err := createTable(db); err != nil {
		log.WithError(err).Error("error creating migration table")
		return err
	}

	completed, err := selectCompleted(db)
	if err != nil && err != sql.ErrNoRows {
		log.WithError(err).Error("error getting completed migrations")
		return err
	}

	for _, migration := range migrations {
		if _, ok := completed[migration.name]; ok {
			continue
		}

		log.WithField("migration", migration.name).
			Info("running migration")
		if _, err := db.Exec(migration.stmt); err != nil {
			log.WithError(err).Errorf("error executing migration %s", migration.name)
			return err
		}

		if err := insertMigration(db, migration.name); err != nil {
			log.WithError(err).Error("error inserting migration to migrations table")
			return err
		}

	}
	return nil
}

func createTable(db *sqlx.DB) error {
	_, err := db.Exec(migrationTableCreate)
	return err
}

func insertMigration(db *sqlx.DB, name string) error {
	_, err := db.Exec(migrationInsert, name)
	return err
}

func selectCompleted(db *sqlx.DB) (map[string]struct{}, error) {
	migrations := map[string]struct{}{}
	rows, err := db.Query(migrationSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		migrations[name] = struct{}{}
	}
	return migrations, nil
}

var (
	migrationTableCreate = `CREATE TABLE IF NOT EXISTS migrations (name VARCHAR(255), UNIQUE(name))`
	migrationInsert      = `INSERT INTO migrations (name) VALUES ($1)`
	migrationSelect      = `SELECT name FROM migrations`
)

var createBuyerTable = `
CREATE TABLE IF NOT EXISTS buyer (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	external_id BIGINT,
	nickname VARCHAR(128),
	email VARCHAR(128),
	first_name VARCHAR(128),
	last_name VARCHAR(128),
	phone_area_code INT,
	phone_number VARCHAR(9),
	billing_info_doc_type VARCHAR(4),
	billing_info_doc_number VARCHAR(14)
)
`

var createReceiverAddressTable = `
CREATE TABLE IF NOT EXISTS receiver_address (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	external_id BIGINT,
	address_line VARCHAR(256),
	street_name  VARCHAR(256),
	street_number VARCHAR(7),
	comment VARCHAR(512),
	zip_code VARCHAR(8),
	city VARCHAR(128),
	state VARCHAR(128),
	country_id VARCHAR(2),
	country_name VARCHAR(128),
	neighborhood_id INT,
	neighborhood_name VARCHAR(128),
	latitude DECIMAL,
	longitude DECIMAL,
	receiver_phone VARCHAR(15)
)
`

var createShippingTable = `
CREATE TABLE IF NOT EXISTS shipping (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	receiver_address_id BIGINT,
	external_id BIGINT,
	shipment_type VARCHAR(128),
	date_created TIMESTAMPTZ,

	CONSTRAINT fk_receiver_address FOREIGN KEY(receiver_address_id) REFERENCES receiver_address(id)
)
`

var createOrdersTable = `
CREATE TABLE IF NOT EXISTS orders (
	id BIGSERIAL NOT NULL PRIMARY KEY,
    external_id BIGINT,
    store_id BIGINT,
    date_created TIMESTAMPTZ,
    date_closed TIMESTAMPTZ,
    last_updated TIMESTAMPTZ,
    total_amount DECIMAL,
    total_shipping DECIMAL,
    total_amount_with_shipping DECIMAL,
    paid_amount DECIMAL,
    expiration_date TIMESTAMPTZ,
    status VARCHAR(128),
    shipping_id BIGINT NOT NULL,
	buyer_id BIGINT NOT NULL,
	
	CONSTRAINT fk_shipping FOREIGN KEY(shipping_id) REFERENCES shipping(id),
	CONSTRAINT fk_buyer FOREIGN KEY(buyer_id) REFERENCES buyer(id)

)
`

var createOrderItemTable = `
CREATE TABLE IF NOT EXISTS orders_item (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	order_id BIGINT NOT NULL,
	external_id VARCHAR(128),
	title VARCHAR(256),
	quantity INT,
	unit_price DECIMAL,
	full_unit_price DECIMAL,

	CONSTRAINT fk_orders FOREIGN KEY(order_id) REFERENCES orders(id)
)
`

var createPaymentTable = `
CREATE TABLE IF NOT EXISTS payment (
	id BIGSERIAL NOT NULL PRIMARY KEY,
    external_id BIGINT,
    external_order_id BIGINT,
    payer_id BIGINT,
    installments INT,
    payment_type VARCHAR(128),
    status VARCHAR(128),
    transaction_amount DECIMAL,
    taxes_amount DECIMAL,
    shipping_cost DECIMAL,
    total_paid_amount DECIMAL,
    installment_amount DECIMAL,
    date_approved TIMESTAMPTZ,
    date_created TIMESTAMPTZ,
	order_id BIGINT NOT NULL,
	
	CONSTRAINT fk_orders FOREIGN KEY(order_id) REFERENCES orders(id)
)
`
