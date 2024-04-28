package http

import "net/http"

type RegisterFunction func(router string, handler func(http.ResponseWriter, *http.Request))
