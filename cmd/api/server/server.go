package server

import (
	"net/http"

	"github.com/apex/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/wesleimp/dc-test/store"
)

// Start server
func Start(c *cli.Context) error {
	addr := c.String("addr")
	conn := c.String("connection-string")

	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowHeaders = []string{
		"Authorization",
		"Content-Type",
		"Access-Control-Allow-Origin",
	}

	db, err := store.Connect(conn)
	if err != nil {
		return errors.Wrap(err, "error creating database connection")
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery(), cors.New(cfg))

	setupRoutes(router, db)

	log.WithField("addr", addr).Info("starting server")
	return http.ListenAndServe(addr, router)
}
