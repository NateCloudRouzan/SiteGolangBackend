package main

import (
    "html/template"
    "net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func form1Handler(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "form1.html")
}

func handlingForm(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("form1.html"))
    if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
    }

    details := ContactDetails{
        Email:   r.FormValue("email"),
        Subject: r.FormValue("subject"),
        Message: r.FormValue("message"),
    }

    // do something with details
    _ = details

    tmpl.Execute(w, struct{ Success bool }{true})    
}

func form2Handler(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "form2.html")
}

func form_2(w http.ResponseWriter, r *http.Request){
       tmpl := template.Must(template.ParseFiles("form2.html"))
    if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
    }
    
    c := LoginInfo{
        Success: true, 
        Fname: r.FormValue("fname"), 
        Lname: r.FormValue("lname"),
        Pword: r.FormValue("pword"),
    }
    
    if c.Fname == "Nate"{
        c.Authorized = true
    }
    
    tmpl.Execute(w, c)   
}

func form3Handler(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "form3.html")
}

func form_3(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("form3.html"))
    tmpl.Execute(w, r.FormValue("photo"))        
}

func form_3_redir(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w, r.Method)
}