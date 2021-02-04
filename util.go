package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"strings"
)
func getHTMLPage(url string)*goquery.Document{
	f, e := os.Open(url)
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()
	doc,err := goquery.NewDocumentFromReader(f)
	if err!=nil{
		return nil
	}
	return doc
}
func checkError(err error) {
	if err != nil {
		print("Error: " + err.Error())
		log.Println(err)
	}
}

//func main() {
//	methods:=newMethods()
//	doc:=getHTMLPage("C:/Users/Admin/Downloads/spam-voice/callspamapi/vn/mobifone/bigdata/callspamapi/utils/StringUtil.html")
//	methods.getAllMethodInformation(doc,"StringUtil")
//
//}

func (methods *Methods) getAllMethodInformation(doc *goquery.Document,className string){
	//var wg sync.WaitGroup
	checkHaveMethod := doc.Find(".method-summary .summary-table").Text()
	packageName:=doc.Find(".sub-title a").Text()
	if checkHaveMethod!=""{
		doc.Find(".summary-table tbody tr").Each(func(i int, s *goquery.Selection){
			typeReturn:=s.Find(".col-first code").Text()
			methodName:= s.Find(".col-second code").Text()
			desc:=s.Find(".col-last div").Text()
			if methodName!=""{
				method:=Method{
					Description: strings.TrimSpace(desc),
					MethodName: methodName,
					TypeReturn: typeReturn,
				}
				methods.Package = packageName
				methods.ClassTitle = className
				//fmt.Println(method)

				methods.List =append(methods.List,method)
			}

		})

	}
	if checkHaveMethod==""{
		method:=Method{
			//ClassTitle: className,
			//Package: packageName,
			Description: "",
			MethodName: "",
			TypeReturn: "",
		}
		methods.Package = packageName
		methods.ClassTitle = className
		methods.List =append(methods.List,method)
	}
}
