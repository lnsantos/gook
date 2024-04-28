package main

import (
	"fmt"
	core "gobook/core/http"
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
	settings := core.CoreHttpCustomHandler{Mux: handler, Logger: logger}

	// register endpoint service
	customHandler := settings.HandleServiceEndpoint(transport.GetBookContract())

	server := &http.Server{
		Addr:           ":8080",
		Handler:        customHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		ErrorLog:       logger,
	}

	log.Fatal(server.ListenAndServe())
}
