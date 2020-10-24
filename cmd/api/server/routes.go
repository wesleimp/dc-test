package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	orderHandler "github.com/wesleimp/dc-test/cmd/api/handlers/order"
	"github.com/wesleimp/dc-test/pkg/buyer"
	"github.com/wesleimp/dc-test/pkg/order"
	"github.com/wesleimp/dc-test/pkg/payment"
	"github.com/wesleimp/dc-test/pkg/shipping"
)

func setupRoutes(r *gin.Engine, db *sqlx.DB) {
	orderStore := order.NewStore(db)
	buyerStore := buyer.NewStore(db)
	paymentStore := payment.NewStore(db)
	shippingStore := shipping.NewStore(db)

	orderService := order.NewService(db, orderStore, buyerStore, paymentStore, shippingStore)

	r.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "health")
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("/orders", orderHandler.Post(orderService))
	}
}
