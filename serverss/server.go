package server

import (
	"io"
	"log"
	"net/http"
)

func Main_func() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/about", about)
	http.HandleFunc("/password", password_reset)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func home_page(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome to my http server home page")
}

func about(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hey the about page of my server")
}

func password_reset(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Reset your password here please")
}
