package main

import (
	"flag"
	"log"
	"net/http"
)

type application struct{}

func main() {
	addr := flag.String("addr", ":4000", "The port address")

	app := &application{}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting server on port %s", *addr)
	err := srv.ListenAndServe()
	log.Fatalln(err)
}
