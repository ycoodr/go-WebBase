package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("The method is " + r.Method)
	fmt.Fprintln(rw, "hello world baby")
}

func pageNf(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "page doesn't work" /*404*/ /*http.StatusNotFound*/, http.StatusConflict)
}

func sayhi(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())

	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	fmt.Fprintf(rw, "Hi, %s your age is %s !", name, age)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/page", pageNf)
	mux.HandleFunc("/error", error)
	mux.HandleFunc("/sayhi", sayhi)

	// http.HandleFunc("/", hello)
	// http.HandleFunc("/page", pageNf)
	// http.HandleFunc("/error", error)
	// http.HandleFunc("/sayhi", sayhi)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	fmt.Println("Server on 3000 port")
	// log.Fatal(http.ListenAndServe("localhost:3000", mux))
	log.Fatal(server.ListenAndServe())
}
