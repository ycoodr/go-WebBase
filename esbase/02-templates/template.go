package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// type Users struct {
// 	Username string
// 	Age      int
// 	Active   bool
// 	Admin    bool
// 	Courses  []Course
// }

type Users struct {
	Username string
	Age      int
}

// type Course struct {
// 	Name string
// }

// func sayhi(name string) string {
// 	return "hi " + name + " from function"
// }

// var templates = template.Must(template.New("T").ParseFiles("index.html", "base.html"))
var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {

	err := templates.ExecuteTemplate(rw, name, data)

	if err != nil {
		// panic(err)
		//http.Error(rw, "can't render template", http.StatusInternalServerError)
		handleError(rw, http.StatusInternalServerError)
	}
}

func handleError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)
}

func index(rw http.ResponseWriter, r *http.Request) {
	// c1 := Course{"Go"}
	// c2 := Course{"Python"}
	// c3 := Course{"Java"}
	// c4 := Course{"Javascript"}

	// fmt.Fprintln(rw, "hello world")
	// template, err := template.ParseFiles("index.html", "base.html")
	// template := template.Must(template.New("index.html").ParseFiles("index.html", "base.html"))
	// functions := template.FuncMap{
	// 	"sayhi": sayhi,
	// }

	// template, err := template.New("index.html").Funcs(functions).ParseFiles("index.html")

	// courses := []Course{c1, c2, c3, c4}
	// user := Users{"Yan", 37, true, false, courses}
	user := Users{"Yan", 37}
	renderTemplate(rw, "index.html", user)
	// template.Execute(rw /*nil*/, user)

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	template.Execute(rw /*nil*/, user)
	// }
}

func register(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "register.html", nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/register", register)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("Server on 3000 port")
	// log.Fatal(http.ListenAndServe("localhost:3000", mux))
	log.Fatal(server.ListenAndServe())
}
