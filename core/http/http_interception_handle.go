package http

import (
	"gobook/core/http_api"
)

type InterceptionNext struct {
	Next bool
}

type InterceptionHandle interface {
	Handle(
		response http_api.Response,
		request *http_api.Request,
	) InterceptionNext
}

type CoreInterceptionHandle struct {
	InnerHandler InterceptionHandle
}
