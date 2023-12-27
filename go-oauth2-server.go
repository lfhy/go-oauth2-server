package main

import (
	"log"
	"os"

	"go-oauth2-server/cmd"
	"go-oauth2-server/config"

	"github.com/urfave/cli"
)

var (
	cliApp        *cli.App
	configBackend string
)

func init() {
	// Initialise a CLI app
	cliApp = cli.NewApp()
	cliApp.Name = "go-oauth2-server"
	cliApp.Usage = "Go OAuth 2.0 Server"
	cliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "config.backend",
			Value:       "file",
			Destination: &configBackend,
			Usage:       "配置文件存储引擎",
		},
		cli.StringFlag{
			Name:        "config.path",
			Value:       "config/go_oauth2_server.json",
			Destination: &config.ConfigPath,
			Usage:       "配置文件读取路径",
		},
	}
}

func main() {
	// Set the CLI app commands
	cliApp.Commands = []cli.Command{
		{
			Name:  "migrate",
			Usage: "run migrations",
			Action: func(c *cli.Context) error {
				return cmd.Migrate(config.ConfigBackendType(configBackend))
			},
		},
		{
			Name:  "loaddata",
			Usage: "load data from fixture",
			Action: func(c *cli.Context) error {
				return cmd.LoadData(c.Args(), config.ConfigBackendType(configBackend))
			},
		},
		{
			Name:  "runserver",
			Usage: "run web server",
			Action: func(c *cli.Context) error {
				return cmd.RunServer(config.ConfigBackendType(configBackend))
			},
		},
	}

	// Run the CLI app
	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
