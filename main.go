package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Welcome struct {
	Name string
	Time time.Time
}

func main() {
	http.HandleFunc("/", serveTemplate)

	fmt.Println("Listening Server...");
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	welcome := Welcome{"Guest User", time.Now()}
	tmpl, _ := template.ParseFiles("index.html")

	year, month, day := welcome.Time.Date()
	fmt.Println(year, month, day)
	_ = tmpl.Execute(w, welcome)
}
