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
				Name:  "load_maps_from_json",
				Usage: "import maps into database from a json file",
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
