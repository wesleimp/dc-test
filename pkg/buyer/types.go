package buyer

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type (
	// Buyer type
	Buyer struct {
		ID                   int64
		ExternalID           int64
		Nickname             string
		Email                string
		PhoneAreaCode        int
		PhoneNumber          string
		FirstName            string
		LastName             string
		BillingInfoDocType   string
		BillingInfoDocNumber string
	}

	// Store interface
	Store interface {
		TxCreate(context.Context, *sqlx.Tx, *Buyer) error
	}
)
