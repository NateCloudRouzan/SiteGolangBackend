package main

import (
	"net/http"
//    "html/template"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir(".")))
    
//    tpl = template.Must(template.ParseFiles("templatePractice.gohtml"))
    
//    err := tpl.Execute(os.Stdout, "From GolangFile")
//    if err != nil{
//        log.Fatlln(err)
//    }
}