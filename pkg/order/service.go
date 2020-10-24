package order

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/wesleimp/dc-test/pkg/buyer"
	"github.com/wesleimp/dc-test/pkg/payment"
	"github.com/wesleimp/dc-test/pkg/shipping"
)

type orderService struct {
	db            *sqlx.DB
	orderStore    Store
	buyerStore    buyer.Store
	paymentStore  payment.Store
	shippingStore shipping.Store
}

// NewService creates a new store service
func NewService(db *sqlx.DB, store Store, bStore buyer.Store, pStore payment.Store, sStore shipping.Store) Service {
	return &orderService{
		db:            db,
		orderStore:    store,
		buyerStore:    bStore,
		paymentStore:  pStore,
		shippingStore: sStore,
	}
}

func (s *orderService) Create(ctx context.Context, order *Order) error {
	tx := s.db.MustBegin()

	if err := s.buyerStore.TxCreate(ctx, tx, order.Buyer); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "error on insert buyer")
	}

	if err := s.shippingStore.TxCreate(ctx, tx, order.Shipping); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "error on insert shipping")
	}

	if err := s.orderStore.TxCreate(ctx, tx, order); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "error on insert order")
	}

	if err := s.paymentStore.TxCreate(ctx, tx, order.Payments, order.ID); err != nil {
		tx.Rollback()
		return errors.Wrap(err, "error on insert payments")
	}

	return tx.Commit()
}
