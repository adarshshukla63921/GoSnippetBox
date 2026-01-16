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
	app := &application{
		logInfo: logInfo,
		logError: logError,
	}
	logInfo.Print("server is runnig at",*addr)
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: logError,
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	logError.Fatal(err)
}