package main

import (
	"gobook/transport"
	"net/http"
	"time"
)

func main() {
	handler := http.NewServeMux()

	handler.HandleFunc(transport.GetBookContract())

	server := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
