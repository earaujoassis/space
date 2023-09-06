package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/tasks"
	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/utils"
)

func init() {
	defer utils.RecoverHandler()
	err := godotenv.Load()
	if err != nil {
		logs.Propagate(logs.Info, "The environment file (.env) doesn't exist; skipping .env")
	} else {
		logs.Propagate(logs.Info, "Application has found a .env file")
	}
	config.LoadConfig()
	cfg := config.GetGlobalConfig()
	logs.Setup(logs.Options{
		Environment: config.Environment(),
		SentryUrl:   cfg.SentryUrl,
	})
}

func main() {
	defer utils.RecoverHandler()
	app := cli.NewApp()
	app.Name = "space"
	app.Version = "0.2.0"
	app.Usage = "A user management microservice; OAuth 2 provider"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "Run actions against the database",
			Subcommands: []cli.Command{
				{
					Name:  "migrate",
					Usage: "Apply migrations to the database upward",
					Flags: []cli.Flag {
						cli.StringFlag{
							Name:  "path",
							Value: "./configs/migrations",
							Usage: "Migrations folder relative path",
						},
					},
					Action: func(c *cli.Context) error {
						tasks.RunMigrations(c.String("path"))
						return nil
					},
				},
				{
					Name:  "rollback",
					Usage: "Apply migrations to the database downward",
					Flags: []cli.Flag {
						cli.StringFlag{
							Name:  "path",
							Value: "./configs/migrations",
							Usage: "Migrations folder relative path",
						},
					},
					Action: func(c *cli.Context) error {
						tasks.RollbackMigrations(c.String("path"))
						return nil
					},
				},
			},
		},
		{
			Name:    "serve",
			Usage:   "Serve the application server",
			Action: func(c *cli.Context) error {
				tasks.Server()
				return nil
			},
		},
		{
			Name:    "launch",
			Usage:   "Apply migrations and serve the application server",
			Flags: []cli.Flag {
				cli.StringFlag{
					Name:  "path",
					Value: "./configs/migrations",
					Usage: "Migrations folder relative path",
				},
			},
			Action: func(c *cli.Context) error {
				tasks.RunMigrations(c.String("path"))
				tasks.Server()
				return nil
			},
		},
		{
			Name:    "client",
			Usage:   "Manage client applications",
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
			Usage:   "Toggle feature flags ON/OFF",
			Action: func(c *cli.Context) error {
				tasks.ToggleFeature()
				return nil
			},
		},
	}

	app.Run(os.Args)
}
