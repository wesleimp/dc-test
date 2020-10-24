package order

import (
	"context"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/wesleimp/dc-test/pkg/buyer"
	"github.com/wesleimp/dc-test/pkg/shipping"
	"github.com/wesleimp/dc-test/store/dbtest"

	_ "github.com/lib/pq"
)

var ctx = context.TODO()

type IntegrationStoreSuite struct {
	suite.Suite

	db    *sqlx.DB
	store Store
}

func (s *IntegrationStoreSuite) SetupTest() {
	s.db = dbtest.Setup("orders")
	s.store = NewStore(s.db)

	_, err := s.db.Exec(stmtInsertBuyerTest)
	assert.NoError(s.T(), err)

	_, err = s.db.Exec(stmtInsertReceiverAddressTest)
	assert.NoError(s.T(), err)

	_, err = s.db.Exec(stmtInsertShippingTest)
	assert.NoError(s.T(), err)
}

func (s *IntegrationStoreSuite) TearDownTest() {
	dbtest.TearDown(s.db, "orders")
}

func (s *IntegrationStoreSuite) TestInsertOrderWithOneItem() {
	assert := assert.New(s.T())

	order := &Order{
		ExternalID:              9987071,
		StoreID:                 282,
		DateCreated:             time.Now().AddDate(0, -1, 0),
		DateClosed:              time.Now(),
		LastUpdated:             time.Now().AddDate(0, 0, -1),
		TotalAmount:             49.9,
		TotalShipping:           5.14,
		TotalAmountWithShipping: 55.04,
		PaidAmount:              55.04,
		ExpirationDate:          time.Now(),
		Items: []*Item{
			{
				ExternalID:    "IT4801901403",
				Title:         "Produto de Testes",
				Quantity:      1,
				UnitPrice:     49.9,
				FullUnitPrice: 49.9,
			},
		},
		Buyer: &buyer.Buyer{
			ID: 1,
		},
		Shipping: &shipping.Shipping{
			ID: 1,
		},
	}

	tx, err := s.db.Beginx()
	assert.NoError(err)

	err = s.store.TxCreate(ctx, tx, order)
	assert.NoError(err)
	assert.NoError(tx.Commit())
	assert.NotZero(order.ID)
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationStoreSuite))
}

var stmtInsertReceiverAddressTest = `
INSERT INTO receiver_address (id) VALUES (1)
`

var stmtInsertShippingTest = `
INSERT INTO shipping (id, receiver_address_id) VALUES (1, 1)
`

var stmtInsertBuyerTest = `
INSERT INTO buyer (id) VALUES (1)
`
