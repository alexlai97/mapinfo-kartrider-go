package main

import (
	"log"
	"os"

	"github.com/alexlai97/mapinfo-kartrider/database"
	"github.com/alexlai97/mapinfo-kartrider/maps"
	"github.com/alexlai97/mapinfo-kartrider/routing"
	"github.com/urfave/cli/v2"
)

func main() {
	// cli
	app := &cli.App{
		Name: "mapinfo",
		Commands: []*cli.Command{
			{
				// TODO: add parameter to the filename
				Name:  "load_default_maps",
				Usage: "accept a json file and load it into database",
				Action: func(c *cli.Context) error {
					maps.InsertDefaultMapsToDB()
					return nil
				},
			},
			{
				// TODO: add parameter to the port
				Name:  "serve",
				Usage: "serve the app",
				Action: func(c *cli.Context) error {
					database.InitDB()
					routing.InitRoutingScheme()
					routing.ServeRouter(":8080")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("app.Run failed", err.Error())
	}
}
