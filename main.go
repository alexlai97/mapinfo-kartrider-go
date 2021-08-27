package main

import (
	"log"
	"os"

	"github.com/alexlai97/mapinfo-kartrider/common"
	"github.com/alexlai97/mapinfo-kartrider/model"
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
					model.InsertDefaultMapsToDB()
					return nil
				},
			},
			{
				// TODO: add parameter to the port
				Name:  "serve",
				Usage: "serve the app",
				Action: func(c *cli.Context) error {
					model.InitDB()
					common.InitRoutingScheme()
					common.ServeRouter(":8080")
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
