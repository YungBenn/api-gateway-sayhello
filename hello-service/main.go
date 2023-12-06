package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")	
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home!")	
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/world", helloWorldHandler)
	http.ListenAndServe(":4000", nil)
}