package main

import (
	"net/http"
    "fmt"
    "html/template"
	"io/ioutil"
    "os"
	"path/filepath"
//    "time"
    "github.com/satori/go.uuid"
)

type User struct{
    username string
    fname string
    lname string
    email string
    birthYear int
    birthMonth int
    birthDay int
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
    cookie, err := r.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
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
}

func udemyProjectsHandler(w http.ResponseWriter, r *http.Request) { //Should be a repo of all of my projects

}

func iconHandler(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "cloudIcon.ico")
}

func studentSealHandler(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "StudentSeal.html")
}


func FileUploadTemplate(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w , r , "file_submit_template.html")
}

func FileUploadHandler(w http.ResponseWriter, req *http.Request) {

	var s string
    tmpl := template.Must(template.ParseFiles("file_submit_template.html"))

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
    
    tmpl.Execute(w, s)
}

func SaveOnServer(w http.ResponseWriter, req *http.Request) {

	var s string
    tmpl := template.Must(template.ParseFiles("file_submit_template.html"))
    
	if req.Method == http.MethodPost {

		// open
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		// store on server
		dst, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, s)
}

func seeUUID(w http.ResponseWriter, req *http.Request){
    cookie, _ := req.Cookie("session")
    a := `Your UUID is: ` + cookie.Value
    fmt.Fprint(w, a)
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

    http.HandleFunc("/redirect301", Redirect301Handler)
    http.HandleFunc("/redirect303", Redirect303Handler)
    http.HandleFunc("/Redirect_303", Redirect303)
    http.HandleFunc("/redirect307", Redirect307Handler)
    http.HandleFunc("/Redirect_307", Redirect307)
    http.HandleFunc("/redir_end", ShowMethod)
    
    http.HandleFunc("/file_submit_template.html", FileUploadTemplate)
    http.HandleFunc("/fileUpload", FileUploadHandler)
    http.HandleFunc("/fileUpload2", SaveOnServer)

    http.HandleFunc("/cookieIncrement", cookieCounter)
    http.HandleFunc("/cookieThrottle", cookieThrottle)
    http.HandleFunc("/cookieUUID", seeUUID)


    http.HandleFunc("/smth/", smthHandler)
    

}