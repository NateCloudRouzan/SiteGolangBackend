package main

import (
	"net/http"
    "fmt"
//    "html/template"
)




func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" && r.URL.Path != "/index.html" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
//    http.FileServer(http.Dir("."))
    http.ServeFile(w , r , "index.html")
    
//    tmpl, err := template.Must(template.ParseFiles("index.html"))

//    err = tmpl.Execute(w, 8)
    
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

func cssServer(w http.ResponseWriter, r *http.Request){ 
//    http.ServeFile(w , r , "/css/main.css")
//    http.ServeFile(w , r , "/css/mq_800-pus.css")
//    http.ServeFile(w , r , "/css/mcleod-reset.css")
    
      http.Handle("/css/", http.FileServer(http.Dir(".")))
}

func handler(w http.ResponseWriter, r *http.Request){ 
    fmt.Fprintf(w, "Path: %s, Length:", r.URL.Path[1:], len(r.URL.Path[1:]))
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

func projectHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w , r , "projects.html")
}

func udemyHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, r.URL.Path)
    
    //http.ServeFile(w , r , "../GolangPractice/UdemyHome.html")
}

func imageHandler(w http.ResponseWriter, r *http.Request){
//    fmt.Fprint(w, r.URL.Path)
    http.ServeFile(w , r, r.URL.Path)

    
 //   http.Handle("/media/img/", http.FileServer(http.Dir(".")))
}

func templateHandler(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("layout.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
}

func init() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/projects.html", projectHandler)
    http.HandleFunc("/main.css", mainCSSHandler)
    http.HandleFunc("/mq_800-plus.css", plusCSSHandler)
    http.HandleFunc("/mcleod-reset.css", resetCSSHandler)

    http.HandleFunc("/css/", cssServer)
    
    http.Handle("/media/img/", http.StripPrefix("/media/img/", http.FileServer(http.Dir("./media/img/"))))
    http.Handle("/GolangPractice/", http.StripPrefix("/GolangPractice/", http.FileServer(http.Dir("./GolangPractice/"))))


    //I think strip prefix is the answer
    
//    http.HandleFunc("/media/img/", imageHandler)
    
  
    //tmpl := template.Must(template.ParseFiles("layout.html"))
//    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		data := TodoPageData{
//			PageTitle: "My TODO list",
//			Todos: []Todo{
//				{Title: "Task 1", Done: false},
//				{Title: "Task 2", Done: true},
//				{Title: "Task 3", Done: true},
//			},
//		}
//		tmpl.Execute(w, data)
//	})
    
    
    http.HandleFunc("/GolangPractice/", udemyHandler)

    

    http.HandleFunc("/smth/", smthHandler)
    
    
    
    
//    tpl = template.Must(template.ParseFiles("templatePractice.gohtml"))
    
//    err := tpl.Execute(os.Stdout, "From GolangFile")
//    if err != nil{
//        log.Fatlln(err)
//    }
}