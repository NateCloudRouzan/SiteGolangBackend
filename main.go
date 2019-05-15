package main

import (
	"net/http"
    "fmt"
    "html/template"
    "io"
	"io/ioutil"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

type LoginInfo struct {
	Success bool
    Authorized bool
    Fname string
	Lname string
    Pword string
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
        http.ServeFile(w , r , "errorPage.html")
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

func errorPageHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w , r , "errorPage.css")
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

func iconHandler(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "cloudIcon.ico")
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

func studentSealHandler(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "StudentSeal.html")
}

func RedirectHandler(w http.ResponseWriter, r *http.Request){
    http.Redirect(w, r, "https://www.youtube.com/watch?v=TOUrLn1FFCA", 301)
}

func FileUploadHandler(w http.ResponseWriter, req *http.Request) {

	var s string
	if req.Method == http.MethodPost {

		// open
		f, _, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// for your information
//		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		// read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}


func init() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
    http.HandleFunc("/", homeHandler)
    
    http.HandleFunc("/cloudIcon.ico", iconHandler)
    http.HandleFunc("/errorPage.html", errorPageHandler)
    http.HandleFunc("/projects.html", projectHandler)
    http.HandleFunc("/StudentSeal.html", studentSealHandler)

    
    http.HandleFunc("/main.css", mainCSSHandler)
    http.HandleFunc("/mq_800-plus.css", plusCSSHandler)
    http.HandleFunc("/mcleod-reset.css", resetCSSHandler)
    http.Handle("/media/img/", http.StripPrefix("/media/img/", http.FileServer(http.Dir("./media/img/"))))

    http.HandleFunc("/GolangPractice", udemyHandler)
 //   http.HandleFunc("/GolangPractice/", udemyProjectsHandler)
    
    
    http.HandleFunc("/GolangPractice/string_template", simpleTemplateString)
    http.HandleFunc("/GolangPractice/int_template", simpleTemplateInt)
    http.HandleFunc("/GolangPractice/slice_template", templateslice)
    http.HandleFunc("/GolangPractice/struct_template", templateStruct)
    http.HandleFunc("/GolangPractice/map_template", templateMap)

    http.HandleFunc("/template_simple.html", template1Layout)
    http.HandleFunc("/template_struct.html", template2Layout)
    http.HandleFunc("/template_slice.html", template3Layout)
    http.HandleFunc("/template_map.html", template4Layout)

    http.HandleFunc("/form1.html", form1Handler)
    http.HandleFunc("/form2.html", form2Handler)
    http.HandleFunc("/form3.html", form3Handler)

    
    http.HandleFunc("/first_form/", handlingForm)
    http.HandleFunc("/second_form/", form_2)
    http.HandleFunc("/third_form", form_3)
    http.HandleFunc("/third_form/", form_3_redir)

    http.HandleFunc("/redirect", RedirectHandler)
    http.HandleFunc("/fileUpload", FileUploadHandler)


    http.HandleFunc("/smth/", smthHandler)
    

}