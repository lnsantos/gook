package transport

import (
	"gobook/core/interceptions"
	"gobook/transport/book"
	"net/http"
)

func GetBookContract() (
	router string,
	handler func(response http.ResponseWriter, request *http.Request),
	excludesInterception []string,
) {
	return "/v1/book", book.Get, []string{interceptions.InterceptionLogger}
}
