package shipping

import (
	"context"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

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
	s.db = dbtest.Setup("shipping")
	s.store = NewStore(s.db)
}

func (s *IntegrationStoreSuite) TearDownTest() {
	dbtest.TearDown(s.db, "shipping")
}

func (s *IntegrationStoreSuite) TestInsertShipping() {
	assert := assert.New(s.T())

	shipping := &Shipping{
		ExternalID:   1,
		ShipmentType: "shipping",
		DateCreated:  time.Now(),
		ReceiverAddress: &ReceiverAddress{
			ExternalID:       3,
			AddressLine:      "Rua Fake de Testes 3454",
			StreetName:       "Rua Fake de Testes",
			StreetNumber:     "3454",
			Comment:          "teste",
			ZipCode:          "85045020",
			City:             "Cidade de Testes",
			State:            "São Paulo",
			CountryID:        "BR",
			CountryName:      "Brasil",
			NeighborhoodID:   nil,
			NeighborhoodName: "Vila de Testes",
			Latitude:         -23.629037,
			Longitude:        -46.712689,
			ReceiverPhone:    "41999999999",
		},
	}

	tx, err := s.db.Beginx()
	assert.NoError(err)

	err = s.store.TxCreate(ctx, tx, shipping)
	if err != nil {
		assert.NoError(tx.Rollback())
	} else {
		assert.NoError(tx.Commit())
	}
	assert.NoError(err)
	assert.NotZero(shipping.ID)
	assert.NotZero(shipping.ReceiverAddress.ID)

}

func TestIntegrationsStoreSuite(t *testing.T) {
	suite.Run(t, new(IntegrationStoreSuite))
}
