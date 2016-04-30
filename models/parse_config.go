package models

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/iyannik0215/CrawlerIShadowsocks/conf"
)

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func ReadProfile() []byte {
	if !Exist(conf.ExecPath + conf.FileName) {
		if err := ioutil.WriteFile(conf.ExecPath+conf.FileName, conf.DefaultConfig, 0664); err != nil {
			log.Fatal("写入文件失败.")
		}
	}
	src, err := ioutil.ReadFile(conf.ExecPath + conf.FileName)
	if err != nil {
		log.Fatal("[ERROR]:", err)
	}
	return src
}

func AnalysisProfile(src []byte) *simplejson.Json {
	json, err := simplejson.NewJson(src)
	if err != nil {
		log.Fatal("[ERROR]:", err)
	}

	return json
}

func ModifyProfile(configs *simplejson.Json, param []interface{}) []byte {
	configs.Del("configs")
	configs.Set("configs", param)
	src, err := configs.EncodePretty()
	if err != nil {
		log.Fatal("[ERROR:]", err)
	}
	// log.Println(string(src))
	return src
}

func WriteProfile(src []byte) {
	err := ioutil.WriteFile(conf.ExecPath+conf.FileName, src, 0644)
	if err != nil {
		log.Fatal("[ERROR]:", err)
	}
}
