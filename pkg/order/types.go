package order

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/wesleimp/dc-test/pkg/buyer"
	"github.com/wesleimp/dc-test/pkg/payment"
	"github.com/wesleimp/dc-test/pkg/shipping"
)

type (
	// Order type
	Order struct {
		ID                      int64
		ExternalID              int64
		StoreID                 int64
		DateCreated             time.Time
		DateClosed              time.Time
		LastUpdated             time.Time
		TotalAmount             float32
		TotalShipping           float32
		TotalAmountWithShipping float32
		PaidAmount              float32
		ExpirationDate          time.Time
		Status                  string
		Buyer                   *buyer.Buyer
		Shipping                *shipping.Shipping
		Payments                []*payment.Payment
		Items                   []*Item
	}

	// Item struct
	Item struct {
		ExternalID    string
		Title         string
		Quantity      int
		UnitPrice     float32
		FullUnitPrice float32
	}

	// Store interface
	Store interface {
		TxCreate(context.Context, *sqlx.Tx, *Order) error
	}

	// Service interface
	Service interface {
		Create(context.Context, *Order) error
	}
)
