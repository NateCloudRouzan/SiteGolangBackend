package main

import (
	"net/http"
    "fmt"
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

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        fmt.Fprint(w, `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Title</title></head><body><h1>Signup</h1>
    <form method="POST" action="/signup">
    Username<br><input type="text" name="username" value="NastyNate"> <br>
    First Name<br><input type="text" name="fname" value="Nate"> <br>
    Last Name<br><input type="text" name="lname" value="Cloud"><br>
    Email<br><input type="text" name="email" value="example@example.com"><br>
    <label for="pass">Password (8 characters minimum):</label>
    <input type="password" id="pass" name="password" minlength="8" required><br>
    <label for="pass">Re-type password:</label>
    <input type="password" id="pass" name="password" minlength="8" required><br>
    <input type="submit" value="Sign Up!"></form></body></html>`)
        return
    }
    
    fmt.Fprint(w, "In post")


    
    //Create account 
    //Link it to that session
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    //take in login info
    //make sure data matches map
    //Link session to username
    
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
    //Delete session cookie 
    //Remove
}

func LoggedInHome(w http.ResponseWriter, r *http.Request) {
    //Show Name username and email
    //Make a birthday countdown
    //

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

    http.HandleFunc("/signup", SignUpHandler)
    http.HandleFunc("/login", LoginHandler)
    http.HandleFunc("/logout", LogoutHandler)
    http.HandleFunc("/account_home", LoggedInHome)
    
    http.HandleFunc("/main.css", mainCSSHandler)
    http.HandleFunc("/mq_800-plus.css", plusCSSHandler)
    http.HandleFunc("/mcleod-reset.css", resetCSSHandler)
    http.Handle("/media/img/", http.StripPrefix("/media/img/", http.FileServer(http.Dir("./media/img/"))))

    
    http.HandleFunc("/GolangPractice", udemyHandler)    
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