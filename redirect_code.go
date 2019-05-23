package main

import (
	"net/http"
    "fmt"
    "time"
)

func RedirectTempHandler(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "redirect_template.html")
}

func Redirect301Handler(w http.ResponseWriter, r *http.Request){
    time.Sleep(10 * time.Second)
    http.Redirect(w, r, "https://www.youtube.com/watch?v=TOUrLn1FFCA", 301)
}

func Redirect303Handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, `<!DOCTYPE html> <html lang="en"><head><meta charset="UTF-8"><title>Title</title></head><body>
<h1>This data will be submitted as a POST method</h1>
<form method="POST" action="/Redirect_303">
    <input type="text" name="fname">
    <input type="submit">
</form></body></html>`)
}

func Redirect303(w http.ResponseWriter, r *http.Request){
    http.Redirect(w, r, "/redir_end", http.StatusSeeOther)
}

func Redirect307Handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, `<!DOCTYPE html> <html lang="en"><head><meta charset="UTF-8"><title>Title</title></head><body>
<h1>This data will be submitted as a POST method</h1>
<form method="POST" action="/Redirect_307">
    <input type="text" name="fname">
    <input type="submit">
</form></body></html>`)
}

func Redirect307(w http.ResponseWriter, r *http.Request){
    http.Redirect(w, r, "/redir_end", http.StatusTemporaryRedirect)
}


func ShowMethod(w http.ResponseWriter, r *http.Request){
    a := "Original method POST, Now the method is: " + r.Method
    fmt.Fprint(w, a)
}