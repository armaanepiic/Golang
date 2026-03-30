package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am Arman. I am software engineer")
}

func main() {
	mux := http.NewServeMux() // router

	mux.HandleFunc("/hello", helloHandler) // route

	mux.HandleFunc("/about", aboutHandler) // route

	fmt.Println("Server running on :3000")

	err := http.ListenAndServe(":3000", mux) // nil

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}