package main

import (
	"net/http"
	"fmt"
)

func init() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("Whattup")
}
