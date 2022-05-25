package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	az "jackminchin.me/dda/az"
	js "jackminchin.me/dda/jupyter_server"
)

func main() {
	app := &cli.App{
		Name:  "dda",
		Usage: "Provision, deploy and used cloud data analytics services",
		Authors: []*cli.Author{
			{
				Name:  "Jack Minchin",
				Email: "jminchin@oxfordeconomics.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "deploy",
				Subcommands: []*cli.Command{
					{
						Name: "jupyter-server",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "resource-group",
								Usage: "The resource group to deploy to",
							},
							&cli.StringFlag{
								Name:     "cores",
								Usage:    "Number of cores to provision",
								Value:    "1",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "memory",
								Usage:    "Amount of memory to provision",
								Value:    "1",
								Required: true,
							},
						},
						Action: js.DeployJupyterServer,
					},
				},
			},
		},
		Before: func(c *cli.Context) error {
			isInstalled := az.IsInstalled()
			if !isInstalled {
				log.Fatal("You must have the az cli installed to use this tool")
			}

			// Check if user is logged in
			isLoggedIn := az.IsLoggedIn()
			if !isLoggedIn {
				log.Fatal("You must be logged in to use this tool. Run \"az login\" first.")
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
