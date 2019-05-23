package main

import (
    "html/template"
    "net/http"
)

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