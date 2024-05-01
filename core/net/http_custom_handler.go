package net

import (
	"encoding/json"
	"gobook/core/interception"
	"gobook/core/netapi"
	"log"
	"net/http"
)

// CoreHttpCustomHandler Struct to init setting with HandleServiceEndpoint /*
type CoreHttpCustomHandler struct {
	Mux    *http.ServeMux
	Logger *log.Logger
}

// interceptionRegister
// Register the interception to the request

// R *net.DefaultError
//   - nil all interception process with successful
//   - not nil some interception break in validation /*
func interceptionRegister(
	origen string,
	writer http.ResponseWriter,
	request *http.Request,
	logger *log.Logger,
	excludes []string,
) *netapi.DefaultError {
	logger.Println("Registering interception for %v", origen)

	excludeLogger := false
	var performError *netapi.DefaultError

	for _, exclude := range excludes {
		if exclude == interception.InterceptionLogger {
			excludeLogger = true
			continue
		}
	}

	if excludeLogger == false && performError == nil {
		performError = interception.LoggerStart().Middleware(writer, request)
	}

	return performError
}

// HandleServiceEndpoint
// register endpoint with perform interception
// before process resource /*
func (httpHandler CoreHttpCustomHandler) HandleServiceEndpoint(
	router string,
	handler func(response http.ResponseWriter, request *http.Request),
	excludes []string,
) {
	logger := httpHandler.Logger
	httpHandler.Mux.HandleFunc(router, func(writer http.ResponseWriter, request *http.Request) {
		if err := interceptionRegister(router, writer, request, logger, excludes); err != nil {

			// Register all header error by interception failed
			for key, value := range err.Headers {
				writer.Header().Set(key, value)
			}

			// Setup status code and serialize struct to response error
			writer.WriteHeader(err.StatusCode)
			data, _ := json.Marshal(err.Data)

			if _, err := writer.Write(data); err != nil {
				logger.Println(err)
				return
			}
		} else {
			// Execute resource endpoint request by client
			handler(writer, request)
		}
	})
}
