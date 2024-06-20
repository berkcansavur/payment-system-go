package server

import (
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
	srv.ListenAndServe()
}
