package server

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Start(r http.Handler) {
	logger := log.New(os.Stdout, "SERVER: ", log.Ldate|log.Ltime|log.Lshortfile)

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      r,
		ErrorLog:     logger,
	}
	log.Fatal(srv.ListenAndServe())
	srv.ListenAndServe()
}
