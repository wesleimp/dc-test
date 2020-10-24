package order

import (
	"github.com/wesleimp/dc-test/pkg/buyer"
	"github.com/wesleimp/dc-test/pkg/order"
	"github.com/wesleimp/dc-test/pkg/payment"
	"github.com/wesleimp/dc-test/pkg/shipping"
)

func toOrder(b body) *order.Order {
	parsedOrder := &order.Order{
		ExternalID:              b.ID,
		StoreID:                 b.StoreID,
		DateCreated:             b.DateCreated,
		DateClosed:              b.DateClosed,
		LastUpdated:             b.LastUpdated,
		TotalAmount:             b.TotalAmount,
		TotalShipping:           b.TotalShipping,
		TotalAmountWithShipping: b.TotalAmountWithShipping,
		PaidAmount:              b.PaidAmount,
		ExpirationDate:          b.ExpirationDate,
		Shipping: &shipping.Shipping{
			ExternalID:   b.Shipping.ID,
			DateCreated:  b.Shipping.DateCreated.UTC(),
			ShipmentType: b.Shipping.ShipmentType,
			ReceiverAddress: &shipping.ReceiverAddress{
				ExternalID:       b.Shipping.ReceiverAddress.ID,
				AddressLine:      b.Shipping.ReceiverAddress.AddressLine,
				City:             b.Shipping.ReceiverAddress.City.Name,
				Comment:          b.Shipping.ReceiverAddress.Comment,
				CountryID:        b.Shipping.ReceiverAddress.Country.ID,
				CountryName:      b.Shipping.ReceiverAddress.Country.Name,
				Latitude:         b.Shipping.ReceiverAddress.Latitude,
				Longitude:        b.Shipping.ReceiverAddress.Longitude,
				NeighborhoodID:   b.Shipping.ReceiverAddress.Neighborhood.ID,
				NeighborhoodName: b.Shipping.ReceiverAddress.Neighborhood.Name,
				ReceiverPhone:    b.Shipping.ReceiverAddress.ReceiverPhone,
				State:            b.Shipping.ReceiverAddress.State.Name,
				StreetName:       b.Shipping.ReceiverAddress.StreetName,
				StreetNumber:     b.Shipping.ReceiverAddress.StreetNumber,
				ZipCode:          b.Shipping.ReceiverAddress.ZipCode,
			},
		},
		Status: b.Status,
		Buyer: &buyer.Buyer{
			ExternalID:           b.Buyer.ID,
			BillingInfoDocNumber: b.Buyer.BillingInfo.DocNumber,
			BillingInfoDocType:   b.Buyer.BillingInfo.DocType,
			Email:                b.Buyer.Email,
			FirstName:            b.Buyer.FirstName,
			LastName:             b.Buyer.LastName,
			Nickname:             b.Buyer.Nickname,
			PhoneAreaCode:        b.Buyer.Phone.AreaCode,
			PhoneNumber:          b.Buyer.Phone.Number,
		},
	}

	items := []*order.Item{}
	for _, item := range b.Items {
		items = append(items, &order.Item{
			ExternalID:    item.Item.ID,
			Title:         item.Item.Title,
			FullUnitPrice: item.FullUnitPrice,
			Quantity:      item.Quantity,
			UnitPrice:     item.UnitPrice,
		})
	}
	parsedOrder.Items = items

	payments := []*payment.Payment{}
	for _, pmt := range b.Payments {
		payments = append(payments, &payment.Payment{
			ExternalID:        pmt.ID,
			DateApproved:      pmt.DateApproved,
			DateCreated:       pmt.DateCreated,
			ExternalOrderID:   pmt.OrderID,
			InstallmentAmount: pmt.InstallmentAmount,
			Installments:      pmt.Installments,
			PayerID:           pmt.PayerID,
			PaymentType:       pmt.PaymentType,
			ShippingCost:      pmt.ShippingCost,
			Status:            pmt.Status,
			TaxesAmount:       pmt.TaxesAmount,
			TotalPaidAmount:   pmt.TotalPaidAmount,
			TransactionAmount: pmt.TransactionAmount,
		})
	}
	parsedOrder.Payments = payments

	return parsedOrder
}
