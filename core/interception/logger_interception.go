package interception

import (
	"gobook/core/netapi"
	"log"
	"net/http"
)

func LoggerStart() InterceptionFunction {
	return InterceptionFunction{
		Middleware: func(response http.ResponseWriter, request *http.Request) *netapi.DefaultError {
			log.Println("Request URL: ", request.URL)

			if request.Header.Get("Authorization") == "" {
				response.Header().Set("Content-Type", "application/json")
				response.WriteHeader(http.StatusUnauthorized)

				log.Printf("Request dropped interception canceled by logger :: %v\n", request.URL)

				return &netapi.DefaultError{
					Headers:    map[string]string{},
					StatusCode: http.StatusUnauthorized,
					Data:       netapi.DefaultResponse{Data: "Unauthorized"},
				}
			}

			return nil
		},
	}
}
