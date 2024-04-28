package http

import (
	"gobook/core/interceptions"
	"log"
	"net/http"
)

type CoreHttpCustomHandler struct {
	Mux    *http.ServeMux
	Logger *log.Logger
}

// interceptionRegister
// Register the interception to the request  /*
func interceptionRegister(
	origen string,
	next http.Handler,
	logger *log.Logger,
	excludes []string,
) http.Handler {
	logger.Println("Registering interception")

	excludeLogger := false

	for _, exclude := range excludes {
		if exclude == interceptions.InterceptionLogger {
			excludeLogger = true
			continue
		}
	}

	if excludeLogger == false {
		next = interceptions.LoggerStart().Middleware(next)
	} else {
		logger.Println("Interception Logger is excluded for %v", origen)
	}

	return next
}

// HandleServiceEndpoint
// register endpoint with perform interception
// before process resource /*
func (httpHandler CoreHttpCustomHandler) HandleServiceEndpoint(
	router string,
	handler func(response http.ResponseWriter, request *http.Request),
	excludes []string,
) http.Handler {
	logger := httpHandler.Logger
	httpHandler.Mux.HandleFunc(router, handler)
	return interceptionRegister(
		router,
		httpHandler.Mux,
		logger,
		excludes,
	)
}
