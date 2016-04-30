package models

import (
	"log"
	"reflect"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Config struct {
	Remarks     string
	Server      string
	Server_port string
	Password    string
	Method      string
}

func GetPage() *goquery.Document {
	for {
		doc, err := goquery.NewDocument("http://www.ishadowsocks.net")
		if err != nil {
			log.Println("[ERROR]:", err)
		}
		return doc
	}
}

func ParsePage(doc *goquery.Document) [3]Config {
	var configArr = [3]Config{}
	count := 0

	doc.Find(".col-lg-4.text-center").EachWithBreak(func(i int, contentSelection *goquery.Selection) bool {
		contentSelection.Find("h4").EachWithBreak(func(i int, configSelection *goquery.Selection) bool {
			text := configSelection.Text()
			result := strings.Split(text, ":")

			switch i {
			case 0:
				configArr[count].Remarks = result[0]
				configArr[count].Server = result[1]
				break
			case 1:
				configArr[count].Server_port = result[1]
				break
			case 2:
				configArr[count].Password = result[1]
				break
			case 3:
				configArr[count].Method = result[1]
				break
			}

			if i == 3 {
				count++
				return false
			}
			return true
		})

		if i == 2 {
			return false
		}
		return true
	})

	return configArr
}

func ConvertMap(conf [3]Config) []interface{} {
	mapSlice := make([]map[string]interface{}, 3)
	o := make([]interface{}, 3)

	for i, cf := range conf {
		tmpMap := make(map[string]interface{})
		elem := reflect.ValueOf(&cf).Elem()
		type_ := elem.Type()

		for j := 0; j < type_.NumField(); j++ {
			tmpMap[strings.ToLower(type_.Field(j).Name)] = elem.Field(j).Interface()
		}

		mapSlice[i] = tmpMap
		o[i] = mapSlice[i]
	}
	return o
}
