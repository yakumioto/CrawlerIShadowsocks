package main

import "log"

func main() {
	configPath := "./gui-config.json"
	iConfig := ParsePage(GetPage())
	WriteProfile(configPath, ModifyProfile(AnalysisProfile(ReadProfile(configPath)), ConvertMap(iConfig)))
	log.Println("执行成功。")
	for {
	}
}
