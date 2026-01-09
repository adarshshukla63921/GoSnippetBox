package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w,r)
		return
	}
	ts , err := template.ParseFiles("ui\\html\\home.html")

	if(err!=nil){
		log.Print(err.Error())
		http.Error(w,"something went wrong",http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w,nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w,"something went wrong rendering the page",http.StatusInternalServerError)
		return
	}
}


func snippetCreate(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w,"Wrong Method Used", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("This will allow you to create snippets"))
}

func snippetView(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil{
		http.NotFound(w,r)
		return
	}

	fmt.Fprintf(w,"id : %d",id)
}