package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "mapinfo",
		Commands: []*cli.Command{
			{
				Name:  "load_default_maps",
				Usage: "accept a json file and load it into database",
				Action: func(c *cli.Context) error {
					insertDefaultMapsToDB()
					return nil
				},
			},
			{
				Name:  "serve",
				Usage: "serve the app",
				Action: func(c *cli.Context) error {
					initDB()
					initRoutingScheme()
					serveRouter(":8080")
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
