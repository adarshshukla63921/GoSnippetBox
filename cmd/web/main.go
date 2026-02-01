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
	// our flag
	addr := flag.String("addr",":8080","HTTP Network Address")
	flag.Parse()
	// params needed for application struct
	logInfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	logError := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// a instance of our application struct
	app := &application{
		logInfo: logInfo,
		logError: logError,
	}
	// this line... is just there to see that our logger is working
	logInfo.Print("server is runnig at",*addr)
	// creating a server with our routes and error logger
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: logError,
		Handler: app.routes(),
	}
	// launch the server
	err := srv.ListenAndServe()
	// something goes wrong, exit main.
	logError.Fatal(err)
}