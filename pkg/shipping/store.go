package shipping

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type shippingStore struct {
	db *sqlx.DB
}

// NewStore creates a new store instance
func NewStore(db *sqlx.DB) Store {
	return &shippingStore{db}
}

func (s *shippingStore) TxCreate(ctx context.Context, tx *sqlx.Tx, shipping *Shipping) error {
	err := insertReceiverAddress(tx, shipping.ReceiverAddress)
	if err != nil {
		return err
	}

	params := toParamsShipping(shipping)
	rows, err := tx.NamedQuery(stmtInsertShipping, params)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&shipping.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func insertReceiverAddress(tx *sqlx.Tx, address *ReceiverAddress) error {
	addressParams := toParamsReceiverAddress(address)

	rows, err := tx.NamedQuery(stmtInsertReceiverAddress, addressParams)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&address.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

var stmtInsertReceiverAddress = `
INSERT INTO receiver_address (
	external_id,
	address_line,
	street_name,
	street_number,
	comment,
	zip_code,
	city,
	state,
	country_id,
	country_name,
	neighborhood_id,
	neighborhood_name,
	latitude,
	longitude,
	receiver_phone
) VALUES (
	:external_id,
	:address_line,
	:street_name,
	:street_number,
	:comment,
	:zip_code,
	:city,
	:state,
	:country_id,
	:country_name,
	:neighborhood_id,
	:neighborhood_name,
	:latitude,
	:longitude,
	:receiver_phone
) RETURNING id
`

var stmtInsertShipping = `
INSERT INTO shipping (
	external_id,
	shipment_type,
	date_created,
	receiver_address_id
) VALUES (
	:external_id,
	:shipment_type,
	:date_created,
	:receiver_address_id
) RETURNING id
`

func toParamsReceiverAddress(address *ReceiverAddress) map[string]interface{} {
	return map[string]interface{}{
		"external_id":       address.ExternalID,
		"address_line":      address.AddressLine,
		"street_name":       address.StreetName,
		"street_number":     address.StreetNumber,
		"comment":           address.Comment,
		"zip_code":          address.ZipCode,
		"city":              address.City,
		"state":             address.State,
		"country_id":        address.CountryID,
		"country_name":      address.CountryName,
		"neighborhood_id":   address.NeighborhoodID,
		"neighborhood_name": address.NeighborhoodName,
		"latitude":          address.Latitude,
		"longitude":         address.Longitude,
		"receiver_phone":    address.ReceiverPhone,
	}
}

func toParamsShipping(s *Shipping) map[string]interface{} {
	return map[string]interface{}{
		"external_id":         s.ExternalID,
		"shipment_type":       s.ShipmentType,
		"date_created":        s.DateCreated,
		"receiver_address_id": s.ReceiverAddress.ID,
	}
}
