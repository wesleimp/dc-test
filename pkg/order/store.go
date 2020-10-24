package order

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type orderStore struct {
	db *sqlx.DB
}

// NewStore creates a new store instance
func NewStore(db *sqlx.DB) Store {
	return &orderStore{db}
}

func (s *orderStore) TxCreate(ctx context.Context, tx *sqlx.Tx, order *Order) error {
	params := toParams(order)
	rows, err := tx.NamedQuery(stmtInsertOrder, params)
	if err != nil {
		return err
	}

	if rows.Next() {
		err := rows.Scan(&order.ID)
		if err != nil {
			return err
		}
	}
	rows.Close()

	return createItems(tx, order)
}

func createItems(tx *sqlx.Tx, order *Order) error {
	for _, item := range order.Items {
		params := toParamsItems(item, order.ID)
		_, err := tx.NamedExec(stmtInsertItem, params)
		return err
	}

	return nil
}

var stmtInsertOrder = `
INSERT INTO orders (
	external_id,
	store_id,
	date_created,
	date_closed,
	last_updated,
	total_amount,
	total_shipping,
	total_amount_with_shipping,
	paid_amount,
	expiration_date,
	status,
	buyer_id,
	shipping_id
) VALUES (
	:external_id,
	:store_id,
	:date_created,
	:date_closed,
	:last_updated,
	:total_amount,
	:total_shipping,
	:total_amount_with_shipping,
	:paid_amount,
	:expiration_date,
	:status,
	:buyer_id,
	:shipping_id
) RETURNING id
`

var stmtInsertItem = `
INSERT INTO orders_item (
	external_id,
	title,
	quantity,
	unit_price,
	full_unit_price,
	order_id
) VALUES (
	:external_id,
	:title,
	:quantity,
	:unit_price,
	:full_unit_price,
	:order_id
)
`

func toParams(o *Order) map[string]interface{} {
	return map[string]interface{}{
		"external_id":                o.ExternalID,
		"store_id":                   o.StoreID,
		"date_created":               o.DateCreated,
		"date_closed":                o.DateClosed,
		"last_updated":               o.LastUpdated,
		"total_amount":               o.TotalAmount,
		"total_shipping":             o.TotalShipping,
		"total_amount_with_shipping": o.TotalAmountWithShipping,
		"paid_amount":                o.PaidAmount,
		"expiration_date":            o.ExpirationDate,
		"status":                     o.Status,
		"buyer_id":                   o.Buyer.ID,
		"shipping_id":                o.Shipping.ID,
	}
}

func toParamsItems(i *Item, orderID int64) map[string]interface{} {
	return map[string]interface{}{
		"external_id":     i.ExternalID,
		"title":           i.Title,
		"quantity":        i.Quantity,
		"unit_price":      i.UnitPrice,
		"full_unit_price": i.FullUnitPrice,
		"order_id":        orderID,
	}
}
