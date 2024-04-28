package interceptions

import "net/http"

type InterceptionFunction struct {
	Middleware func(next http.Handler) http.Handler
}

const (
	InterceptionLogger = "logger"
)
