package payment

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type paymentStore struct {
	db *sqlx.DB
}

// NewStore creates a new store instance
func NewStore(db *sqlx.DB) Store {
	return &paymentStore{db}
}

func (s *paymentStore) TxCreate(ctx context.Context, tx *sqlx.Tx, payments []*Payment, orderID int64) error {
	for _, payment := range payments {
		params := toParams(payment, orderID)
		stmt, args, err := tx.BindNamed(stmtInsertPayment, params)
		if err != nil {
			return err
		}

		rows, err := tx.Query(stmt, args...)
		if err != nil {
			return err
		}
		defer rows.Close()

		if rows.Next() {
			err := rows.Scan(&payment.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

var stmtInsertPayment = `
INSERT INTO payment (
	external_id,
	external_order_id,
	payer_id,
	installments,
	payment_type,
	status,
	transaction_amount,
	taxes_amount,
	shipping_cost,
	total_paid_amount,
	installment_amount,
	date_approved,
	date_created,
	order_id
) VALUES (
	:external_id,
	:external_order_id,
	:payer_id,
	:installments,
	:payment_type,
	:status,
	:transaction_amount,
	:taxes_amount,
	:shipping_cost,
	:total_paid_amount,
	:installment_amount,
	:date_approved,
	:date_created,
	:order_id
) RETURNING id
`

func toParams(p *Payment, orderID int64) map[string]interface{} {
	return map[string]interface{}{
		"external_id":        p.ExternalID,
		"external_order_id":  p.ExternalOrderID,
		"payer_id":           p.PayerID,
		"installments":       p.Installments,
		"payment_type":       p.PaymentType,
		"status":             p.Status,
		"transaction_amount": p.TransactionAmount,
		"taxes_amount":       p.TaxesAmount,
		"shipping_cost":      p.ShippingCost,
		"total_paid_amount":  p.TotalPaidAmount,
		"installment_amount": p.InstallmentAmount,
		"date_approved":      p.DateApproved,
		"date_created":       p.DateCreated,
		"order_id":           orderID,
	}
}
