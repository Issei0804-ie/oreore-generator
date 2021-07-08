package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	html := header()
	f, err := os.OpenFile("test.txt", os.O_RDWR, 777)
	if err != nil {
		log.Fatalln(err.Error())
	}
	all, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	html = html + "\n" + convert(string(all))
	html = html + "\n" + footer()

	fmt.Printf("%v",html)
}

func header()(h string){
	return "<html>\n" +
		"<META HTTP-EQUIV=\"Content-Type\" CONTENT=\"text/html; charset=utf-8\">\n" +
		"<title>issei's home page</title>\n" +
		"<body>\n"
}

func footer()(f string){
	return "</body>\n"+
		"</html>"
}


func convert(b string)(r string){
	rows := strings.Split(b, "\n")
	re1 := regexp.MustCompile(`t:`)
	re2 := regexp.MustCompile(`- `)
	re3 := regexp.MustCompile(`--:s`)
	re4 := regexp.MustCompile(`--:e`)
	for _, row := range rows{
		resultRow := ""
		if re1.MatchString(row){
			resultRow = re1.ReplaceAllString(row, "<h1>") + "</h1>\n"
		}else if re2.MatchString(row){
			resultRow = re2.ReplaceAllString(row, "<LI>") + "</LI>\n"
		}else if re3.MatchString(row){
			resultRow = re3.ReplaceAllString(row, "<ul>") + "\n"
		}else if re4.MatchString(row){
			resultRow = re4.ReplaceAllString(row, "</ul>") + "\n"
		}else{
			resultRow = row + "\n"
		}
		r = r + resultRow
	}
	return
}