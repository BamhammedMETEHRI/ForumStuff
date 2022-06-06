package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/hello", homeHandler)
	http.HandleFunc("/form", formHandler)

	port := "5500"
	fmt.Printf("Starting Server at Port :%v\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
func homeHandler(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(rw, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(rw, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(rw, "Hello ! ! ! !")
}
func formHandler(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(rw, "ParseForm() err %v\n", err)
	}
	fmt.Fprintf(rw, "Post Request Successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(rw, "name : %v\n", name)
	fmt.Fprintf(rw, "address : %v\n", address)
}
