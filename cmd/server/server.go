package server

import (
	"log"
	"net/http"
	"time"
)

func Start(r http.Handler) {
	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      r,
	}
	log.Fatal(srv.ListenAndServe())
	srv.ListenAndServe()
}
