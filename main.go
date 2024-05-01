package main

import (
	"fmt"
	"gobook/core/net"
	"gobook/transport"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Print("Server is running on port 8080\n")
	handler := http.NewServeMux()
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	settings := net.CoreHttpCustomHandler{Mux: handler, Logger: logger}

	// register endpoint service
	settings.HandleServiceEndpoint(transport.GetBookContract())

	server := &http.Server{
		Addr:           ":8081",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		ErrorLog:       logger,
	}

	log.Fatal(server.ListenAndServe())
}
