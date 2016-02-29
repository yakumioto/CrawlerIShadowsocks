package main

import (
	"io/ioutil"
	"log"

	"github.com/bitly/go-simplejson"
)

func ReadProfile(path string) []byte {
	src, err := ioutil.ReadFile(path)
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

	// configs, ok := json.CheckGet("configs")
	// if !ok {
	// 	log.Fatal("[ERROR]:", err)
	// }

	return json
}

func ModifyProfile(configs *simplejson.Json, param []interface{}) []byte {
	configs.Del("configs")
	configs.Set("configs", param)
	src, err := configs.MarshalJSON()
	if err != nil {
		log.Fatal("[ERROR:]", err)
	}
	// log.Println(string(src))
	return src
}

func WriteProfile(path string, src []byte) {
	err := ioutil.WriteFile(path, src, 0644)
	if err != nil {
		log.Fatal("[ERROR]:", err)
	}
}
