package main

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/urfave/cli/v2"
	"github.com/wesleimp/dc-test/cmd/api/server"

	// tool import
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	log.SetHandler(text.Default)

	app := &cli.App{
		Name:  "api",
		Usage: "",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "addr",
				Usage:   "Application address",
				Value:   ":8005",
				EnvVars: []string{"ADDR"},
			},
			&cli.StringFlag{
				Name:    "connection-string",
				Usage:   "Database connection string",
				Value:   "postgres://postgres:postgres@localhost/dctest?sslmode=disable", // default pg config
				EnvVars: []string{"CONNECTION_STRING"},
			},
		},
		Action: server.Start,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}
