package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(W, "Hello, World!")
	if r.URL.Path != "/hello" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello, World!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() error: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successfull\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Printf("Name = %s \n", name)
	fmt.Printf("Address = %s \n", address)
}
func main() {
	fmt.Println("Hello, World!")
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at :8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
