package main

import (
	"net/http"
    "fmt"
//    "html/template"
)




func homeHandler(w http.ResponseWriter, r *http.Request) {
    
    if r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
    
    tmpl, err := template.New("name").Parse("index.html")
    // Error checking elided
    err = tmpl.Execute(w, "Hello")
    
//    http.FileServer(http.Dir("."))
    
//    fmt.Fprint(w, "welcome home")
}

func smthHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/smth/" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
    fmt.Fprint(w, "welcome smth")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "custom 404")
    }
}

func handler(w http.ResponseWriter, r *http.Request){ 
    fmt.Fprintf(w, "Path: %s, Length:", r.URL.Path[1:], len(r.URL.Path[1:]))
}

func init() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
        http.HandleFunc("/", homeHandler)
    http.HandleFunc("/smth/", smthHandler)
    
//    tpl = template.Must(template.ParseFiles("templatePractice.gohtml"))
    
//    err := tpl.Execute(os.Stdout, "From GolangFile")
//    if err != nil{
//        log.Fatlln(err)
//    }
}