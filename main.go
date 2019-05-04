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
    //    http.Handle("/GolangPractice/", http.StripPrefix("/GolangPractice/", http.FileServer(http.Dir("./GolangPractice/"))))
    if r.URL.Path == "/GolangPractice/template1" {
        templateHandler(w,r)
        return
    }
    
    http.Handle("/GolangPractice/", http.StripPrefix("/GolangPractice/", http.FileServer(http.Dir("./GolangPractice/"))))
    //fmt.Fprint(w, r.URL.Path)
}
                                                         

func templateHandler(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("http://cloudrouzan.com/GolangPractice/layout.html"))
  //  http.HandleFunc("/GolangPractice/template1", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
//	})
}

func init() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/projects.html", projectHandler)
    http.HandleFunc("/main.css", mainCSSHandler)
    http.HandleFunc("/mq_800-plus.css", plusCSSHandler)
    http.HandleFunc("/mcleod-reset.css", resetCSSHandler)
    http.Handle("/media/img/", http.StripPrefix("/media/img/", http.FileServer(http.Dir("./media/img/"))))

    http.HandleFunc("/GolangPractice", udemyHandler)
    http.HandleFunc("/GolangPractice/", udemyProjectsHandler)

    http.HandleFunc("/smth/", smthHandler)
    

}