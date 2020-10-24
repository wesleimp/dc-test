package buyer

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type buyerStore struct {
	db *sqlx.DB
}

// NewStore creates a new instance of buyer store
func NewStore(db *sqlx.DB) Store {
	return &buyerStore{db}
}

func (s *buyerStore) TxCreate(ctx context.Context, tx *sqlx.Tx, buyer *Buyer) error {
	params := toParams(buyer)
	rows, err := tx.NamedQuery(stmtInsert, params)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&buyer.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

var stmtInsert = `
INSERT INTO buyer (
	external_id, 
	nickname, 
	email, 
	first_name, 
	last_name, 
	phone_area_code, 
	phone_number, 
	billing_info_doc_type, 
	billing_info_doc_number
) VALUES (
	:external_id, 
	:nickname, 
	:email, 
	:first_name, 
	:last_name, 
	:phone_area_code, 
	:phone_number, 
	:billing_info_doc_type, 
	:billing_info_doc_number
) RETURNING id
`

func toParams(b *Buyer) map[string]interface{} {
	return map[string]interface{}{
		"external_id":             b.ExternalID,
		"nickname":                b.Nickname,
		"email":                   b.Email,
		"first_name":              b.FirstName,
		"last_name":               b.LastName,
		"phone_area_code":         b.PhoneAreaCode,
		"phone_number":            b.PhoneNumber,
		"billing_info_doc_type":   b.BillingInfoDocType,
		"billing_info_doc_number": b.BillingInfoDocNumber,
	}
}
