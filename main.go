package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sync"
)
func main() {

	var name string = "file_exported"

	//if _, err := os.Stat("./output/"+name+".txt");  err != nil {
	//	if os.IsExist(err) {
	//		fmt.Println("in")
	//		e := os.Remove("./output/"+name+".txt")
	//		if e != nil {
	//			log.Fatal(e)
	//		}
	//	} else {
	//		// other error
	//	}
	//}
	var wg sync.WaitGroup
	file,_:=ioutil.ReadFile("categories.json")
	categoires:=Categories{}
	_ = json.Unmarshal([]byte(file),&categoires)
	jobs := make(chan Category,1000)
	results :=make(chan Methods,1000)
	crawlAllFromCategories(jobs,results,&wg)
	for i:=0;i<len(categoires.List) ;i++{
		jobs<-categoires.List[i]
	}
	close(jobs)
	for i:=0;i<len(categoires.List) ;i++ {
		select {
		case methodsReceive, open := <-results:
			if !open {
				break
			}
			writeToFile(methodsReceive,name)
		}
	}
	wg.Wait()
}
func writeToFile(methodReceive Methods,name string){
	f, _ := os.OpenFile("./output/"+name+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	io.WriteString(f, "Tên Package: "+methodReceive.Package+"\n")
	io.WriteString(f, "Tên Class: "+methodReceive.ClassTitle+"\n")
	io.WriteString(f, "Type Return|"+"Method Name|"+"Description"+"\n")
	for j:=0;j<len(methodReceive.List);j++{
		io.WriteString(f, methodReceive.List[j].TypeReturn+"|"+methodReceive.List[j].MethodName+"|"+methodReceive.List[j].Description+"\n")
	}
	io.WriteString(f, "\n")
}
func crawlAllFromCategories(jobs<- chan Category,results chan <- Methods,wg *sync.WaitGroup){
	for w:=1;w<=10;w++{
		wg.Add(1)
		go worker(w,jobs,results,wg)
	}
}
var mutex sync.Mutex
func worker(id int,jobs<-chan Category,results chan <- Methods,wg *sync.WaitGroup){
	defer wg.Done()
	if _, err := os.Stat("./output"); os.IsNotExist(err) {
		os.Mkdir("./output", 0755)
	}
	for j:= range jobs {//duyệt qua tất cả category
		methods:=newMethods()
		//dt := time.Now()
		fmt.Println("worker: ", id, "processing job: ", j)
		res:=getHTMLPage(j.URL)
		if res == nil{
			fmt.Println("Page not found")
			return
		}
		methods.getAllMethodInformation(res,j.Title)
		results<-*methods
	}
}
