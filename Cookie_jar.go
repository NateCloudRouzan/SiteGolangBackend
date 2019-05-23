package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func cookieCounter(w http.ResponseWriter, req *http.Request){
    cookie, err := req.Cookie("my-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
			Path: "/",
		}
	}

	count, _ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)
    a := `<button onclick="window.location.href = 'https://cloudrouzan.com/cookieIncrement';">Cookie Incrementer</button>` + cookie.Value
	fmt.Fprint(w, a)
}

func cookieThrottle(w http.ResponseWriter, req *http.Request){
    cookie, err := req.Cookie("num-visits")
    timer, not_there := req.Cookie("hold-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "num-visits",
			Value: "0",
			Path: "/",
		}
	}

    if not_there == http.ErrNoCookie {
		timer = &http.Cookie{
			Name:  "hold-cookie",
			Value: "Hold",
			Path: "/",
            MaxAge: 7,
		}
        http.SetCookie(w, timer)

       //Increment cookie if hold isnt there
       count, _ := strconv.Atoi(cookie.Value)
	   count++
	   cookie.Value = strconv.Itoa(count)
       http.SetCookie(w, cookie)
       a := `<button onclick="window.location.href = 'https://cloudrouzan.com/cookieThrottle';">Cookie Incrementer</button>` + cookie.Value
	   fmt.Fprint(w, a)
       return
	}
    
    q := `<button onclick="window.location.href = 'https://cloudrouzan.com/cookieThrottle';">Cookie Incrementer</button>` + cookie.Value + " " + timer.Value
    fmt.Fprint(w, q)
}