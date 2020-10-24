package buyer

import (
	"context"
	"testing"

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
	s.db = dbtest.Setup("buyer")
	s.store = NewStore(s.db)
}

func (s *IntegrationStoreSuite) TearDownTest() {
	dbtest.TearDown(s.db, "buyer")
}

func (s *IntegrationStoreSuite) TestInsertBuyer() {
	assert := assert.New(s.T())

	buyer := &Buyer{
		ExternalID:           1,
		BillingInfoDocType:   "CPF",
		BillingInfoDocNumber: "11111111111",
		Email:                "FOO@test.com",
		FirstName:            "FOO",
		LastName:             "BAR",
		Nickname:             "FOOBAR",
		PhoneAreaCode:        47,
		PhoneNumber:          "123456789",
	}
	tx, err := s.db.Beginx()
	assert.NoError(err)

	err = s.store.TxCreate(ctx, tx, buyer)
	assert.NoError(tx.Commit())
	assert.NoError(err)
	assert.NotZero(buyer.ID)

	rows, err := s.db.Query("SELECT nickname FROM buyer WHERE external_id = 1")
	defer rows.Close()
	assert.NoError(err)

	var nick string
	assert.True(rows.Next())
	assert.NoError(rows.Scan(&nick))
	assert.Equal(nick, buyer.Nickname)
}

func TestIntegrationRepositorySuite(t *testing.T) {
	suite.Run(t, new(IntegrationStoreSuite))
}
