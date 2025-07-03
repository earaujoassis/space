package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"

	"github.com/earaujoassis/space/internal"
	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/tasks"
	"github.com/earaujoassis/space/internal/utils"
)

func main() {
	defer utils.RecoverHandler()
	err := godotenv.Load()
	if err == nil {
		logs.Propagate(logs.LevelInfo, "Application has found a .env file")
	}
	cfg, err := config.Load()
	if err != nil {
		logs.Propagatef(logs.LevelPanic, "Could not load configuration: %s\n", err)
	}
	logs.Setup(logs.Options{
		Environment: cfg.Environment,
		Release:     cfg.Release(),
		SentryUrl:   cfg.SentryUrl,
	})
	app := cli.NewApp()
	app.Name = "space"
	app.Version = internal.Version
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
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "path",
							Value: "./configs/migrations",
							Usage: "Migrations folder relative path",
						},
					},
					Action: func(c *cli.Context) error {
						tasks.RunMigrations(cfg, c.String("path"))
						return nil
					},
				},
				{
					Name:  "rollback",
					Usage: "Apply migrations to the database downward",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "path",
							Value: "./configs/migrations",
							Usage: "Migrations folder relative path",
						},
					},
					Action: func(c *cli.Context) error {
						tasks.RollbackMigrations(cfg, c.String("path"))
						return nil
					},
				},
			},
		},
		{
			Name:  "serve",
			Usage: "Serve the application server",
			Action: func(c *cli.Context) error {
				tasks.Server(cfg)
				return nil
			},
		},
		{
			Name:  "launch",
			Usage: "Apply migrations and serve the application server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "path",
					Value: "./configs/migrations",
					Usage: "Migrations folder relative path",
				},
			},
			Action: func(c *cli.Context) error {
				tasks.RunMigrations(cfg, c.String("path"))
				tasks.Server(cfg)
				return nil
			},
		},
		{
			Name:  "workers",
			Usage: "Start the workers server/service",
			Action: func(c *cli.Context) error {
				tasks.Workers(cfg)
				return nil
			},
		},
		{
			Name:  "scheduler",
			Usage: "Start the scheduler service",
			Action: func(c *cli.Context) error {
				tasks.Scheduler(cfg)
				return nil
			},
		},
		{
			Name:  "feature",
			Usage: "Toggle feature flags ON/OFF",
			Action: func(c *cli.Context) error {
				tasks.ToggleFeature(cfg)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
