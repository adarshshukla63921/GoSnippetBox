package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("ui\\static"))
	mux.Handle("/static/",http.StripPrefix("/static",fileserver))

	
	mux.HandleFunc("/",home)

	mux.HandleFunc("/create",snippetCreate)

	mux.HandleFunc("/view",snippetView)


	err := http.ListenAndServe(":8080",mux)

	log.Fatal(err)
}