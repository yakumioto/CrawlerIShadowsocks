package main

func main() {
	configPath := "./gui-config.json"
	iConfig := ParsePage(GetPage())
	WriteProfile(configPath, ModifyProfile(AnalysisProfile(ReadProfile(configPath)), ConvertMap(iConfig)))
}
