package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)
type application struct{
	logInfo  *log.Logger
	logError *log.Logger
}
func main() {
	addr := flag.String("addr",":8080","HTTP Network Address")
	flag.Parse()
	logInfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	logError := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("ui\\static"))
	mux.Handle("/static/",http.StripPrefix("/static",fileserver))

	app := &application{
		logInfo: logInfo,
		logError: logError,
	}

	mux.HandleFunc("/",app.home)

	mux.HandleFunc("/create",app.snippetCreate)

	mux.HandleFunc("/view",app.snippetView)

	logInfo.Print("server is runnig at",*addr)
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: logError,
		Handler: mux,
	}
	err := srv.ListenAndServe()
	logError.Fatal(err)
}