package book

import (
	"fmt"
	nt "gobook/core/network"
	"net/http"
)
import source "gobook/datasource/books"

func Get(w http.ResponseWriter, r *http.Request) {
	interaction, err := source.Get()

	if err != nil {
		_, err = w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		if err != nil {
			return
		}
	}

	nt.SendRequestHttp{R: w}.SendRequestC(http.StatusOK, interaction)
}
