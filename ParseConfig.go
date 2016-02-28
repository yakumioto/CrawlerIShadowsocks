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

func AnalysisProfile(src []byte) {
	json, err := simplejson.NewJson(src)
	if err != nil {
		log.Fatal("[ERROR]:", err)
	}

	configs, ok := json.CheckGet("configs")
	if !ok {
		log.Println("0")
	}
	arr, err := configs.Array()
	if err != nil {
		log.Fatal("[ERROR]:", err)
	}

	for _, config := range arr {
		log.Println(config)
	}
	//log.Println(config)
}
