package shipping

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	// Shipping type
	Shipping struct {
		ID              int64
		ExternalID      int64
		ShipmentType    string
		DateCreated     time.Time
		ReceiverAddress *ReceiverAddress
	}

	// ReceiverAddress type
	ReceiverAddress struct {
		ID               int64
		ExternalID       int64
		AddressLine      string
		StreetName       string
		StreetNumber     string
		Comment          string
		ZipCode          string
		City             string
		State            string
		CountryID        string
		CountryName      string
		NeighborhoodID   *int
		NeighborhoodName string
		Latitude         float32
		Longitude        float32
		ReceiverPhone    string
	}

	// Store interface
	Store interface {
		TxCreate(context.Context, *sqlx.Tx, *Shipping) error
	}
)
