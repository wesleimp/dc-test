package processor

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wesleimp/dc-test/pkg/order"
)

func toPayload(o *order.Order) *Payload {
	payload := &Payload{
		City:            o.Shipping.ReceiverAddress.City,
		Complement:      o.Shipping.ReceiverAddress.Comment,
		Country:         o.Shipping.ReceiverAddress.CountryID,
		DateOrderCreate: o.DateCreated,
		DeliveryFee:     floatToString(o.TotalShipping),
		District:        o.Shipping.ReceiverAddress.NeighborhoodName,
		ExternalCode:    strconv.FormatInt(o.ExternalID, 10),
		Latitude:        o.Shipping.ReceiverAddress.Latitude,
		Longitude:       o.Shipping.ReceiverAddress.Longitude,
		Number:          o.Shipping.ReceiverAddress.StreetNumber,
		PostalCode:      o.Shipping.ReceiverAddress.ZipCode,
		State:           o.Shipping.ReceiverAddress.State,
		StoreID:         o.StoreID,
		Street:          o.Shipping.ReceiverAddress.StreetName,
		SubTotal:        floatToString(o.TotalAmount),
		Total:           floatToString(o.TotalAmountWithShipping),
		TotalShipping:   floatToString(o.TotalShipping),
		Customer: &Customer{
			ExternalCode: strconv.FormatInt(o.Buyer.ExternalID, 10),
			Contact:      fmt.Sprintf("%v%s", o.Buyer.PhoneAreaCode, o.Buyer.PhoneNumber),
			Email:        o.Buyer.Email,
			Name:         o.Buyer.Nickname,
		},
	}

	items := []*Item{}
	for _, item := range o.Items {
		items = append(items, &Item{
			ExternalCode: item.ExternalID,
			Name:         item.Title,
			Price:        item.UnitPrice,
			Quantity:     item.Quantity,
			SubItems:     []struct{}{},
			Total:        item.FullUnitPrice,
		})
	}
	payload.Items = items

	payments := []*Payment{}
	for _, pmt := range o.Payments {
		payments = append(payments, &Payment{
			Type:  strings.ToUpper(pmt.PaymentType),
			Value: pmt.TotalPaidAmount,
		})
	}
	payload.Payments = payments

	return payload
}

func floatToString(v float32) string {
	return fmt.Sprintf("%.2f", v)
}
