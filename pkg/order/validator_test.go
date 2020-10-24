package order

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wesleimp/dc-test/pkg/buyer"
	"github.com/wesleimp/dc-test/pkg/payment"
	"github.com/wesleimp/dc-test/pkg/shipping"
)

func TestAllFieldsvalidation(t *testing.T) {
	order := &Order{
		Buyer: &buyer.Buyer{},
		Items: []*Item{
			{Title: "", ExternalID: ""},
		},
		Shipping: &shipping.Shipping{
			ReceiverAddress: &shipping.ReceiverAddress{},
		},
		Payments: []*payment.Payment{{
			ID: 1,
		}},
	}

	errs := order.Validate()
	assert.Equal(t, 14, len(errs))

	assert.Contains(t, errs, errInvalidCommentLen)
	assert.Contains(t, errs, errInvalidCityLen)
	assert.Contains(t, errs, errInvalidBuyerNicknameLen)
	assert.Contains(t, errs, errInvalidCountryIDLen)
	assert.Contains(t, errs, errInvalidNeighborhoodNameLen)
	assert.Contains(t, errs, errInvalidStateLen)
	assert.Contains(t, errs, errInvalidStreetNameLen)
	assert.Contains(t, errs, errInvalidStreetNumberLen)
	assert.Contains(t, errs, errInvalidZipCodeLen)
	assert.Contains(t, errs, errInvalidBuyerEmail)
}
