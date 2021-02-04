package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"strings"
)

var URLAllClass string = "C:/Users/Admin/Downloads/spam-voice/callspamapi/allclasses-index.html"
var index = strings.LastIndex(URLAllClass,"/")
var root = URLAllClass[:index+1]
func (categories *Categories) getAllCategories(doc *goquery.Document) {
	fmt.Println(root)
	doc.Find(".summary-table tbody tr td a:first-child").Each(func(i int, s *goquery.Selection){
		cateLink,_:= s.Attr("href")
		cateTitle:=s.Text()
		fmt.Println(cateLink)
		fmt.Println(cateTitle)
		category := Category{
			Title: cateTitle,
			URL: root+cateLink,
		}
		categories.Total++
		categories.List = append(categories.List,category)
	})
}
func crawlAllCategories(){
	categories := newCategoires()
	res := getHTMLPage(URLAllClass)
	categories.getAllCategories(res)
	categoriesJson,err := json.Marshal(categories)
	checkError(err)
	err = ioutil.WriteFile("categories.json",categoriesJson,0644)
	checkError(err)
}
func main() {
	crawlAllCategories()
}
