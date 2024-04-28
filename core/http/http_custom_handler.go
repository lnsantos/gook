package http

import (
	"net/http"
)

type CoreHttpCustomHandler struct {
	Mux           *http.ServeMux
	Interceptions []InterceptionHandle
}

func (httpHandler CoreHttpCustomHandler) HandleServiceEndpoint(
	router string, handler func(http.ResponseWriter, *http.Request),
) {
	httpHandler.Mux.HandleFunc(
		router,
		func(response http.ResponseWriter, request *http.Request) {
			handlerInterception(
				httpHandler.Interceptions,
				response,
				request,
				handler,
			)
		},
	)
}

func handlerInterception(
	middlewares []InterceptionHandle,
	response http.ResponseWriter,
	request *http.Request,
	handler func(http.ResponseWriter, *http.Request),
) {

	for _, middleware := range middlewares {
		_ = middleware.Handle(response, request)
	}

	handler(response, request)
}
