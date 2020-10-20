package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Starting server...")
	http.ListenAndServe(":5000", nil)

}
