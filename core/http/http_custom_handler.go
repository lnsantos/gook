package http

import "net/http"

type CoreHttpCustomHandler struct {
	Mux *http.ServeMux
}

func (mux *http.ServeMux) HandleServiceEndpoint(pattern string, handler func(http.ResponseWriter, *http.Request)) {

}
