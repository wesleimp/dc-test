package processor

import "time"

type (
	// Payload contains the order info
	Payload struct {
		ExternalCode    string     `json:"externalCode"`
		StoreID         int64      `json:"storeId"`
		SubTotal        string     `json:"subTotal"`
		DeliveryFee     string     `json:"deliveryFee"`
		Total           string     `json:"total"`
		Country         string     `json:"country"`
		State           string     `json:"state"`
		City            string     `json:"city"`
		District        string     `json:"district"`
		Street          string     `json:"street"`
		Complement      string     `json:"complement"`
		Latitude        float32    `json:"latitude"`
		Longitude       float32    `json:"longitude"`
		DateOrderCreate time.Time  `json:"dtOrderCreate"`
		PostalCode      string     `json:"postalCode"`
		Number          string     `json:"number"`
		Customer        *Customer  `json:"customer"`
		Items           []*Item    `json:"items"`
		Payments        []*Payment `json:"payments"`
		TotalShipping   string     `json:"total_shipping"`
	}

	// Customer info
	Customer struct {
		ExternalCode string `json:"externalCode"`
		Name         string `json:"name"`
		Email        string `json:"email"`
		Contact      string `json:"contact"`
	}

	// Item info
	Item struct {
		ExternalCode string     `json:"externalCode"`
		Name         string     `json:"name"`
		Price        float32    `json:"price"`
		Quantity     int        `json:"quantity"`
		Total        float32    `json:"total"`
		SubItems     []struct{} `json:"subItems"`
	}

	// Payment info
	Payment struct {
		Type  string  `json:"type"`
		Value float32 `json:"value"`
	}
)
