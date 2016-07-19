package main

import (
    "os"

    "github.com/urfave/cli"

    "github.com/earaujoassis/space/tasks"
)

func main() {
    app := cli.NewApp()
    app.Name = "space"
    app.Usage = "A user management microservice; OAuth 2 provider"

    app.Commands = []cli.Command{
        {
            Name:    "serve",
            Aliases: []string{"s"},
            Usage:   "Serve the application server",
            Action:  func(c *cli.Context) error {
                tasks.Server("./web/public", "web/templates/*.html")
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
    }

    app.Run(os.Args)
}
