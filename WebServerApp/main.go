package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("./*.html"))
}


func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parseform %v", err)
		return
	}
	tpl.ExecuteTemplate(w,"form.html",nil)
	fmt.Fprintf(w, "Post form successful\n")
	name := r.Form.Get("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "address = %s\n", address)
	fmt.Printf("%+v\n", r.Form)

}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not supported", http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(w, "Hello Yogesh ")

}

func main() {
	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloFunc)

	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
