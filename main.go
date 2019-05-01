package main

import (
	"net/http"
    "fmt"
//    "html/template"
)


func handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func init() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
    http.HandleFunc("/", handler)
    
//    tpl = template.Must(template.ParseFiles("templatePractice.gohtml"))
    
//    err := tpl.Execute(os.Stdout, "From GolangFile")
//    if err != nil{
//        log.Fatlln(err)
//    }
}