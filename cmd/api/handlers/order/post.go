package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wesleimp/dc-test/pkg/order"
	"github.com/wesleimp/dc-test/pkg/processor"
)

// Post handler for order
func Post(service order.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var in body
		err := json.NewDecoder(c.Request.Body).Decode(&in)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"errors": []string{fmt.Sprintf("error parsing request body: %s", err)}})
			return
		}

		order := toOrder(in)
		errs := order.Validate()
		if len(errs) > 0 {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"errors": errs})
			return
		}

		err = service.Create(c.Request.Context(), order)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		res, status, err := processor.Process(order)
		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"errors": []string{err.Error()}})
			return
		}
		if status >= 400 {
			c.JSON(status, map[string]interface{}{"errors": []string{res}})
			return
		}

		c.JSON(status, string(res))
	}
}

// Body is a representative payload
type body struct {
	ID                      int64     `json:"id"`
	StoreID                 int64     `json:"store_id"`
	DateCreated             time.Time `json:"date_created"`
	DateClosed              time.Time `json:"date_closed"`
	LastUpdated             time.Time `json:"last_updated"`
	TotalAmount             float32   `json:"total_amount"`
	TotalShipping           float32   `json:"total_shipping"`
	TotalAmountWithShipping float32   `json:"total_amount_with_shipping"`
	PaidAmount              float32   `json:"paid_amount"`
	ExpirationDate          time.Time `json:"expiration_date"`
	Items                   []struct {
		Item struct {
			ID    string `json:"id"`
			Title string `json:"title"`
		} `json:"item"`
		Quantity      int     `json:"quantity"`
		UnitPrice     float32 `json:"unit_price"`
		FullUnitPrice float32 `json:"full_unit_price"`
	} `json:"order_items"`
	Payments []struct {
		ID                int64     `json:"id"`
		OrderID           int64     `json:"order_id"`
		PayerID           int64     `json:"payer_id"`
		Installments      int       `json:"installments"`
		PaymentType       string    `json:"payment_type"`
		Status            string    `json:"status"`
		TransactionAmount float32   `json:"transaction_amount"`
		TaxesAmount       float32   `json:"taxes_amount"`
		ShippingCost      float32   `json:"shipping_amount"`
		TotalPaidAmount   float32   `json:"total_paid_amount"`
		InstallmentAmount float32   `json:"installment_amount"`
		DateApproved      time.Time `json:"date_approved"`
		DateCreated       time.Time `json:"date_created"`
	}
	Shipping struct {
		ID              int64     `json:"id"`
		ShipmentType    string    `json:"shipment_type"`
		DateCreated     time.Time `json:"date_created"`
		ReceiverAddress struct {
			ID           int64  `json:"id"`
			AddressLine  string `json:"address_line"`
			StreetName   string `json:"street_name"`
			StreetNumber string `json:"street_number"`
			Comment      string `json:"comment"`
			ZipCode      string `json:"zip_code"`
			City         struct {
				Name string `json:"name"`
			} `json:"city"`
			State struct {
				Name string `json:"name"`
			} `json:"state"`
			Country struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"country"`
			Neighborhood struct {
				ID   *int   `json:"id"`
				Name string `json:"name"`
			} `json:"neighborhood"`
			Latitude      float32 `json:"latitude"`
			Longitude     float32 `json:"longitude"`
			ReceiverPhone string  `json:"receiver_phone"`
		} `json:"receiver_address"`
	} `json:"shipping"`
	Status string `json:"status"`
	Buyer  struct {
		ID       int64  `json:"id"`
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
		Phone    struct {
			AreaCode int    `json:"area_code"`
			Number   string `json:"number"`
		} `json:"phone"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		BillingInfo struct {
			DocType   string `json:"doc_type"`
			DocNumber string `json:"doc_number"`
		} `json:"billing_info"`
	} `json:"buyer"`
}
