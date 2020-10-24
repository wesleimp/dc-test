package order

import (
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"
)

var (
	errInvalidLatitude                = "Invalid value for Latitude property"
	errInvalidLongitude               = "Invalid value for Longitude property"
	errInvalidCityLen                 = "Invalid length for City property"
	errInvalidCommentLen              = "Invalid length for Comment property"
	errInvalidCountryIDLen            = "Invalid length for CountryID property"
	errInvalidTotalShipping           = "Invalid value for TotalShipping property"
	errInvalidNeighborhoodNameLen     = "Invalid length for NeighborhoodName property"
	errInvalidOrderID                 = "Invalid value for ID property"
	errInvalidStreetNumberLen         = "Inavlid length for StreetNumber property"
	errInvalidZipCodeLen              = "Invalid length for ZipCode property"
	errInvalidStateLen                = "Invalid length for State property"
	errInvalidStoreID                 = "Invalid value for StoreID property"
	errInvalidStreetNameLen           = "Invalid length for StreetName property"
	errInvalidTotalAmount             = "Invalid value for TotalAmount property"
	errInvalidTotalAmountWithShipping = "Invalid value for TotalAmountWithShipping property"
	errInvalidBuyerID                 = "Inavlid value for Buyer.ID property"
	errInvalidBuyerPhone              = "Inavlid value for Buyer.Phone property"
	errInvalidBuyerEmail              = "Inavlid value for Buyer.Email property"
	errInvalidBuyerNicknameLen        = "Inavlid length for Buyer.Nickname property"
	errInavlidItemIDLen               = "Invalid length for Item.ID property"
	errInvalidItemTitleLen            = "Invalid length for Item.Title property"
	errInvalidItemUnitPrice           = "Invalid value for Item.UnitPrice property"
	errInvalidItemQuantity            = "Invalid value for Item.Quantity property"
	errInvalidItemFullUnitPrice       = "Invalid value for Item.FullUnitPrice property"
	errInvalidPaymentType             = "Invalid value for Payment.PaymentType property"
	errInvalidPaymentTotalPaidAmount  = "Invalid value for Payment.TotalPaidAmount property"
)

// Validate order
func (o *Order) Validate() []string {
	eerros := []string{}

	if !govalidator.IsLatitude(fmt.Sprintf("%f", o.Shipping.ReceiverAddress.Latitude)) {
		eerros = append(eerros, errInvalidLatitude)
	}
	if !govalidator.IsLongitude(fmt.Sprintf("%f", o.Shipping.ReceiverAddress.Longitude)) {
		eerros = append(eerros, errInvalidLongitude)
	}
	if !govalidator.IsByteLength(o.Shipping.ReceiverAddress.City, 1, 128) {
		eerros = append(eerros, errInvalidCityLen)
	}
	if !govalidator.IsByteLength(o.Shipping.ReceiverAddress.Comment, 1, 256) {
		eerros = append(eerros, errInvalidCommentLen)
	}
	if !govalidator.IsByteLength(o.Shipping.ReceiverAddress.CountryID, 1, 2) {
		eerros = append(eerros, errInvalidCountryIDLen)
	}
	if !govalidator.IsFloat(fmt.Sprintf("%f", o.TotalShipping)) {
		eerros = append(eerros, errInvalidTotalShipping)
	}
	if !govalidator.IsFloat(fmt.Sprintf("%f", o.TotalAmount)) {
		eerros = append(eerros, errInvalidTotalAmount)
	}
	if !govalidator.IsFloat(fmt.Sprintf("%f", o.TotalAmountWithShipping)) {
		eerros = append(eerros, errInvalidTotalAmountWithShipping)
	}
	if !govalidator.IsByteLength(o.Shipping.ReceiverAddress.NeighborhoodName, 1, 128) {
		eerros = append(eerros, errInvalidNeighborhoodNameLen)
	}
	if !govalidator.IsInt(strconv.FormatInt(o.ExternalID, 10)) {
		eerros = append(eerros, errInvalidOrderID)
	}
	if !govalidator.IsByteLength(o.Shipping.ReceiverAddress.StreetName, 1, 128) {
		eerros = append(eerros, errInvalidStreetNameLen)
	}
	if !govalidator.IsByteLength(o.Shipping.ReceiverAddress.StreetNumber, 1, 7) {
		eerros = append(eerros, errInvalidStreetNumberLen)
	}
	if !govalidator.IsByteLength(o.Shipping.ReceiverAddress.ZipCode, 1, 8) {
		eerros = append(eerros, errInvalidZipCodeLen)
	}
	if !govalidator.IsByteLength(o.Shipping.ReceiverAddress.State, 1, 128) {
		eerros = append(eerros, errInvalidStateLen)
	}
	if !govalidator.IsInt(strconv.FormatInt(o.StoreID, 10)) {
		eerros = append(eerros, errInvalidStoreID)
	}

	// validate buyer info
	if !govalidator.IsInt(strconv.FormatInt(o.Buyer.ExternalID, 10)) {
		eerros = append(eerros, errInvalidBuyerID)
	}
	if !govalidator.IsEmail(o.Buyer.Email) {
		eerros = append(eerros, errInvalidBuyerEmail)
	}
	if !govalidator.IsByteLength(o.Buyer.Nickname, 1, 128) {
		eerros = append(eerros, errInvalidBuyerNicknameLen)
	}
	if !govalidator.IsInt(strconv.Itoa(o.Buyer.PhoneAreaCode)) || !govalidator.IsByteLength(o.Buyer.PhoneNumber, 1, 9) {
		eerros = append(eerros, errInvalidBuyerPhone)
	}

	for i, item := range o.Items {
		if !govalidator.IsByteLength(item.Title, 1, 128) {
			eerros = append(eerros, errorAtIndex(i, errInvalidItemTitleLen))
		}
		if !govalidator.IsFloat(fmt.Sprintf("%f", item.UnitPrice)) {
			eerros = append(eerros, errorAtIndex(i, errInvalidItemUnitPrice))
		}
		if !govalidator.IsFloat(fmt.Sprintf("%f", item.FullUnitPrice)) {
			eerros = append(eerros, errorAtIndex(i, errInvalidItemFullUnitPrice))
		}
		if !govalidator.IsInt(strconv.Itoa(item.Quantity)) {
			eerros = append(eerros, errorAtIndex(i, errInvalidItemQuantity))
		}
		if !govalidator.IsByteLength(item.ExternalID, 1, 128) {
			eerros = append(eerros, errorAtIndex(i, errInavlidItemIDLen))
		}
	}

	for i, pmt := range o.Payments {
		if !govalidator.IsByteLength(pmt.PaymentType, 1, 128) {
			eerros = append(eerros, errorAtIndex(i, errInvalidPaymentType))
		}
		if !govalidator.IsFloat(fmt.Sprintf("%f", pmt.TotalPaidAmount)) {
			eerros = append(eerros, errorAtIndex(i, errInvalidPaymentTotalPaidAmount))
		}
	}

	return eerros
}

func errorAtIndex(index int, err string) string {
	return fmt.Sprintf("%s: index %d", err, index)
}
