package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := "static/" + title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/home.html")
	t.Execute(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8042", nil)
}
