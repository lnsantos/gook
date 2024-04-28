package interceptions

import (
	"gobook/core/network"
	"net/http"
)

type InterceptionFunction struct {
	Middleware func(response http.ResponseWriter, request *http.Request) *network.DefaultError
}

const (
	InterceptionLogger = "logger"
)
