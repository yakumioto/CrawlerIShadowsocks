package conf

var (
	URL = "http://www.ishadowsocks.net/"

	ExecPath = "./"
	FileName = "gui-config.json"

	DefaultConfig = []byte(`{
"configs": [],
"strategy": "com.shadowsocks.strategy.balancing",
"index": -1,
"global": false,
"enabled": true,
"shareOverLan": false,
"isDefault": false,
"localPort": 1080,
"pacUrl": null,
"useOnlinePac": false,
"availabilityStatistics": false,
"autoCheckUpdate": true,
"logViewer": null
}`)
)
