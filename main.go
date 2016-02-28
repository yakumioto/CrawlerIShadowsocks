package main

import "log"

func main() {
	log.Println(ParsePage(GetPage()))
	AnalysisProfile(ReadProfile("testData/gui-config.json"))
}
