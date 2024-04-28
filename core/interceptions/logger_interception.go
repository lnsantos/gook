package interceptions

import (
	"gobook/core/network"
	"log"
	"net/http"
)

func LoggerStart() InterceptionFunction {
	return InterceptionFunction{
		Middleware: func(response http.ResponseWriter, request *http.Request) *network.DefaultError {
			log.Println("Request URL: ", request.URL)

			if request.Header.Get("Authorization") == "" {
				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusUnauthorized)

				log.Printf("Request dropped interception canceled by logger :: %v\n", request.URL)

				return &network.DefaultError{
					Headers:    map[string]string{},
					StatusCode: http.StatusUnauthorized,
					Data:       network.DefaultResponse{Data: "Unauthorized"},
				}
			}

			return nil
		},
	}
}
