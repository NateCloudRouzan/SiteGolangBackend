package main

import (
	"net/http"
    "html/template"
)


type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func simpleTemplateString (w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("template_simple.html"))
    tmpl.Execute(w, `From template_code.go we doin it!`)
}

func simpleTemplateInt (w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("template_simple.html"))
    tmpl.Execute(w, 48)
}

func templateslice (w http.ResponseWriter, r *http.Request){
    greetings :=[]string{"Yo", "Hi", "Howdy", "Hello", "Wassup"}
    tmpl := template.Must(template.ParseFiles("template_slice.html"))
    tmpl.Execute(w, greetings)
}

func templateMap (w http.ResponseWriter, r *http.Request){
    nicknames := map[string]string{
        "Nate": "Diesel",
        "Amanuel": "Manny",
        "Nathan": "Nose",
        "Caelen": "Duece",
       
    }
    tmpl := template.Must(template.ParseFiles("template_map.html"))
    tmpl.Execute(w, nicknames)
    
}

func templateStruct(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("template_struct.html"))
  // fmt.Fprintf(w, "MAde it passed line 1")
     data := TodoPageData{
			PageTitle: "Template Executed Right if this the title",
			Todos: []Todo{
				{Title: "Get shiii done", Done: false},
				{Title: "sent from golang template", Done: true},
				{Title: "This is a good feelin", Done: true},
			},
      }
     tmpl.Execute(w, data)
}

//Need to implement time Package
func template1Layout(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "template_struct.html")
}

func template2Layout(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "template_simple.html")
}

func template4Layout(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "template_map.html")
}

func template3Layout(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "template_slice.html")
}