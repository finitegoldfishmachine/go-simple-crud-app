package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"simplecrudapp/config"
)

func main() {
	c := config.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      mux,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server :8080")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!")
}
