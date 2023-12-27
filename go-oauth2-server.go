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
	port          int
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
			Usage: "数据库初始化或迁移",
			Action: func(c *cli.Context) error {
				return cmd.Migrate(config.ConfigBackendType(configBackend))
			},
		},
		{
			Name:  "loaddata",
			Usage: "从配置文件加载数据",
			Action: func(c *cli.Context) error {
				return cmd.LoadData(c.Args(), config.ConfigBackendType(configBackend))
			},
		},
		{
			Name:  "server",
			Usage: "运行Web服务",
			Action: func(c *cli.Context) error {
				return cmd.RunServer(config.ConfigBackendType(configBackend), port)
			},
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "run.port",
					Value:       8080,
					Destination: &port,
					Usage:       "服务运行端口",
				}},
		},
	}
	// Run the CLI app
	if err := cliApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
