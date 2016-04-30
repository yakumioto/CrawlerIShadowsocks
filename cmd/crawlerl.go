package cmd

import (
	"fmt"
	"os/exec"

	"github.com/codegangsta/cli"
	"github.com/iyannik0215/CrawlerIShadowsocks/conf"
	"github.com/iyannik0215/CrawlerIShadowsocks/models"
)

func RunCrawlerl(ctx *cli.Context) {
	models.WriteProfile(
		models.ModifyProfile(
			models.AnalysisProfile(models.ReadProfile()),
			models.ConvertMap(models.ParsePage(models.GetPage()))))

	fmt.Println("执行完成, 关闭Dos退出.")
	exec.Command(conf.ExecPath + "Shadowsocks.exe").Run()
}
