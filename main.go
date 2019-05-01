package main

import (
	"net/http"
//    "html/template"
)


func handler(){
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