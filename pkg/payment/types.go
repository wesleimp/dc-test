package payment

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	// Payment struct
	Payment struct {
		ID                int64
		ExternalID        int64
		ExternalOrderID   int64
		PayerID           int64
		Installments      int
		PaymentType       string
		Status            string
		TransactionAmount float32
		TaxesAmount       float32
		ShippingCost      float32
		TotalPaidAmount   float32
		InstallmentAmount float32
		DateApproved      time.Time
		DateCreated       time.Time
	}

	// Store interface
	Store interface {
		TxCreate(context.Context, *sqlx.Tx, []*Payment, int64) error
	}
)
