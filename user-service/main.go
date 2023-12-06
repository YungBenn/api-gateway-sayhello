package main

import (
	"fmt"
	"net/http"
)

func welcomeUserHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Welcome User!")
}

func getUserHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Get User!")
}

func main() {
	http.HandleFunc("/", welcomeUserHandler)
	http.HandleFunc("/get", getUserHandler)
	http.ListenAndServe(":4001", nil)	
}