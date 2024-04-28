package http

import "net/http"

type InterceptionNext struct {
	Next bool
}

type InterceptionHandle interface {
	Handle(
		response http.ResponseWriter,
		request *http.Request,
	) InterceptionNext
}
