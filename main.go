package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Config struct {
	Title  string
	Server string
	Port   string
	Passwd string
	Method string
}

func GetPage() *goquery.Document {
	doc, err := goquery.NewDocument("http://www.ishadowsocks.com/")
	if err != nil {
		log.Fatal("[ERROR]:", err)
	}
	return doc
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
				configArr[count].Title = result[0]
				configArr[count].Server = result[1]
				break
			case 1:
				configArr[count].Port = result[1]
				break
			case 2:
				configArr[count].Passwd = result[1]
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

func main() {
	doc := GetPage()
	log.Println(ParsePage(doc))
}
