package interception

import (
	"gobook/core/netapi"
	"net/http"
)

type InterceptionFunction struct {
	Middleware func(response http.ResponseWriter, request *http.Request) *netapi.DefaultError
}

const (
	InterceptionLogger = "logger"
)
