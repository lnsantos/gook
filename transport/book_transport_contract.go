package transport

import (
	"gobook/transport/book"
	"net/http"
)

func GetBookContract() (router string, handler func(response http.ResponseWriter, request *http.Request)) {
	return "/v1/book", book.Get
}
