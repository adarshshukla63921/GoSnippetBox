package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/",home)

	mux.HandleFunc("/create",snippetCreate)

	mux.HandleFunc("/view",snippetView)


	err := http.ListenAndServe(":8080",mux)

	log.Fatal(err)
}