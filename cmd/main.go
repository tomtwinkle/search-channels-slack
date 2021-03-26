package main

import (
	"context"
	"fmt"
	"github.com/tomtwinkle/search-channels-slack/config"
	"github.com/tomtwinkle/search-channels-slack/slacklib"
	"github.com/urfave/cli"
)

var version = "unknown"
var revision = "unknown"

func main() {
	app := cli.NewApp()
	app.Name = "Slack Tool"
	app.Usage = "Slack tool"
	app.Author = "tomtwinkle"
	app.Version = fmt.Sprintf("slack tool version %s.rev-%s", version, revision)
	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "設定ファイルの作成",
			Action: func(c *cli.Context) error {
				cfg := config.NewConfig()
				if _, err := cfg.Init(); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:      "post",
			ShortName: "p",
			Usage:     "メッセージを投稿する",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "message, m",
					Usage: "メッセージ",
					Required: true,
					Value:    "",
				},
				cli.StringFlag{
					Name: "channel, c",
					Usage: "投稿チャンネル",
					Required: true,
					Value:    "",
				},
			},
			Action: func(c *cli.Context) error {
				cf := config.NewConfig()
				cfg, err := cf.Read()
				if err != nil {
					return err
				}
				sc := slacklib.NewSlackClient(cfg)
				switch c.String("type") {
					return cli.ShowSubcommandHelp(c)
				},
			},
		},
		if err := app.Run(os.Args); err != nil{
		log.Fatal(err)
	}
	}
