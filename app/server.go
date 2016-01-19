package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Sensor struct {
	Name  string
	Value float64
}

func Sensorhandler(w http.ResponseWriter, r *http.Request) {
	sensor := Sensor{Name: r.URL.Path[len("/sensor/"):], Value: 42.0}
	fmt.Println(sensor.Name)
	t, _ := template.ParseFiles("static/sensor.html")
	t.Execute(w, sensor)
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/home.html")
	test := "hejsan svejsan"
	t.Execute(w, test)
}

func main() {
	http.HandleFunc("/home/", handler)
	http.HandleFunc("/sensor/", Sensorhandler)
	http.ListenAndServe(":8042", nil)
}
