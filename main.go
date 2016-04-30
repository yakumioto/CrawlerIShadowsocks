package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/iyannik0215/CrawlerIShadowsocks/cmd"
	"github.com/iyannik0215/CrawlerIShadowsocks/conf"
)

func main() {
	app := cli.NewApp()
	app.Name = "Crawlerl shadowoskcs"
	app.Usage = "Crawlerl http://www.ishadowsocks.net free ss."
	app.Version = "2.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path, p",
			Value:       "./",
			Usage:       "Shadowsocks.exe 的路径, 默认路径为 当前路径, 推荐.",
			Destination: &conf.ExecPath,
		},
	}

	app.Action = cmd.RunCrawlerl

	app.Run(os.Args)
}
