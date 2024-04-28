package main

import (
	"fmt"
	core "gobook/core/http"
	"gobook/transport"
	"net/http"
	"time"
)

func registerInterception() []core.InterceptionHandle {
	return []core.InterceptionHandle{}
}

func main() {
	handler := http.NewServeMux()
	settings := core.CoreHttpCustomHandler{
		Mux:           handler,
		Interceptions: registerInterception(),
	}

	settings.HandleServiceEndpoint(transport.GetBookContract())

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

	_ = fmt.Sprint("Server running on port 8080")
}
