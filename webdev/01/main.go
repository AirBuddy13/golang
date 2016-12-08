package main

import (
	"net/http"
	//"fmt"
	"fmt"
	"html/template"
)

var tpl *template.Template
var mux *http.ServeMux

func init() {
	tpl = template.Must(template.ParseGlob("*"))
	mux = http.NewServeMux()
	mux.HandleFunc("/", test)
}

func factorial(n int) chan int {
	c := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		c <- total
		close(c)
	}()

	return c
}

func main() {
	total := factorial(4)
	fmt.Println("Total:", <-total)

	//	http.ListenAndServe(":8081", mux)
}

func test(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
	//fmt.Fprint(w, "Test")
}
