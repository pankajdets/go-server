package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name =%s\n", name)
	fmt.Fprintf(w, "Address =%s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello There! This is Pankaj")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) //We are telling Go to check static directory
	http.Handle("/", fileServer)                        //start handling root route (/) to send to fileServer
	http.HandleFunc("/form", formHandler)               // handle /form route
	http.HandleFunc("/hello", helloHandler)             // handle /hello route

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
