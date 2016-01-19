package main

import (
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/home.html")
	test := "hejsan svejsan"
	t.Execute(w, test)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8042", nil)
}
