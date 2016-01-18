package main

import (
	"html/template"
	"net/http"
)

type Page interface {
	Title string,
	tmpl string
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
	t, _ := template.ParseFiles("static/" + p.tmpl + ".html")
	t.Execute(w, r)

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/home.html")
	test := "hejsan svejsan"
	if test == "a" {
	}
	t.Execute(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8042", nil)
}
