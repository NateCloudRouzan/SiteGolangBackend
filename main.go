package main

import (
	"net/http"
    "fmt"
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

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" && r.URL.Path != "/index.html" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
    http.ServeFile(w , r , "index.html")
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

func projectHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w , r , "projects.html")
}

func mainCSSHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w , r , "main.css")
}
func plusCSSHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w , r , "mq_800-plus.css")
}
func resetCSSHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w , r , "mcleod-reset.css")
}

func udemyHandler(w http.ResponseWriter, r *http.Request) {//Should be a portal for my webpages
    http.ServeFile(w , r , "UdemyHome.html")

//    fmt.Fprint(w, "welcome home")
}

func udemyProjectsHandler(w http.ResponseWriter, r *http.Request) { //Should be a repo of all of my projects

}
    
func simpleTemplateString (w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("template_simple.html"))
    tmpl.Execute(w, `From Main.go we doin it!`)
}

func simpleTemplateInt (w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("template_simple.html"))
    tmpl.Execute(w, 48)
}

func templateslice (w http.ResponseWriter, r *http.Request){
    greetings :=[]string{"Yo", "Hi", "Howdy", "Hello", "Wassup"}
//    tmpl := template.Must(template.ParseFiles("template_simple.html"))
//    tmpl.Execute(w, 48)
}

func templateMap (w http.ResponseWriter, r *http.Request){

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

func template2Layout(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "template_simple.html")
}

func template1Layout(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "template_struct.html")
}

func iconHandler(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "cloudIcon.ico")
}

func init() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/cloudIcon.ico", iconHandler)

    http.HandleFunc("/projects.html", projectHandler)
    http.HandleFunc("/main.css", mainCSSHandler)
    http.HandleFunc("/mq_800-plus.css", plusCSSHandler)
    http.HandleFunc("/mcleod-reset.css", resetCSSHandler)
    http.Handle("/media/img/", http.StripPrefix("/media/img/", http.FileServer(http.Dir("./media/img/"))))

    http.HandleFunc("/GolangPractice", udemyHandler)
    http.HandleFunc("/GolangPractice/", udemyProjectsHandler)
    
    
    http.HandleFunc("/GolangPractice/template1", templateStruct)
    http.HandleFunc("/GolangPractice/template_string", simpleTemplateString)
    http.HandleFunc("/GolangPractice/template_int", simpleTemplateInt)



    http.HandleFunc("/template_struct.html", template2Layout)
    http.HandleFunc("/template_simple.html", template1Layout)



    http.HandleFunc("/smth/", smthHandler)
    

}