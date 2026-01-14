package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/"{
		http.NotFound(w,r)
		return
	}
	files := []string{
		"ui\\html\\base.html",
		"ui\\html\\nav.html",
		"ui\\html\\pages\\home.html",
	}

	ts , err := template.ParseFiles(files...)

	if err != nil {
		app.logError.Print(err)
		http.Error(w,"Something went wrong",http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w,"base",nil)

	if err != nil {
		app.logError.Print(err)
		http.Error(w,"Something went wrong",http.StatusInternalServerError)
		return
	}
}


func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w,"Wrong Method Used", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("This will allow you to create snippets"))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil{
		http.NotFound(w,r)
		return
	}

	fmt.Fprintf(w,"id : %d",id)
}