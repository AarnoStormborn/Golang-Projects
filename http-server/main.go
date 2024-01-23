package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method Not Supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Saying Hello from HelloHandler")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() Error: %v", err)
	}
	fmt.Println("POST Request Successful")

	name := r.FormValue("name")
	age := r.FormValue("age")
	fmt.Fprintf(w, "Name: %v\nAge: %v\n", name, age)
}

func main() {
	fileServer := http.FileServer(http.Dir("./templates"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting Server at Port 8080: ")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
