package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/tasks"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("> The environment file (.env) doesn't exist; skipping .env\n")
	}

	config.LoadConfig()
}

func main() {
	app := cli.NewApp()
	app.Name = "space"
	app.Usage = "A user management microservice; OAuth 2 provider"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Usage:   "Serve the application server",
			Action: func(c *cli.Context) error {
				tasks.Server()
				return nil
			},
		},
		{
			Name:    "client",
			Aliases: []string{"c"},
			Usage:   "Manage client application",
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "Create a new client application",
					Action: func(c *cli.Context) error {
						tasks.CreateClient()
						return nil
					},
				},
			},
		},
		{
			Name:    "feature",
			Aliases: []string{"f"},
			Usage:   "Toggle features flags ON/OFF",
			Action: func(c *cli.Context) error {
				tasks.ToggleFeature()
				return nil
			},
		},
	}

	app.Run(os.Args)
}
