package interceptions

import (
	"log"
	"net/http"
)

func LoggerStart() InterceptionFunction {
	return InterceptionFunction{
		Middleware: func(next http.Handler) http.Handler {

			return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
				log.Println("Request URL: ", request.URL)

				if request.Header.Get("Authorization") == "" {
					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusUnauthorized)

					log.Printf("Request dropped interception canceled by logger :: %v\n", request.URL)
					if _, err := response.Write([]byte(`{"message": "Unauthorized"}`)); err != nil {
						log.Printf("Request dropped interception canceled by logger with error in client none response %v\n", err.Error())
						return
					}
				} else {
					next.ServeHTTP(response, request)
				}
			})
		},
	}
}
