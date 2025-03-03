package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler)
	fmt.Printf("Server is running port 8080!")
	http.ListenAndServe(":8080", r)
}