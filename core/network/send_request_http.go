package network

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	key_time    = "time"
	key_message = "message"
	key_status  = "status"
)

type DefaultResponse struct {
	Message *string `json:"message,omitempty"`
	Data    any     `json:"data"`
	Time    *int32  `json:"time,omitempty"`
	Status  *bool   `json:"status,omitempty"`
}

type SendRequestHttp struct {
	R http.ResponseWriter
}

func (w SendRequestHttp) SendRequestC(
	code int,
	response any,
) {
	w.SendRequest(code, response, nil)
}

func (w SendRequestHttp) SendRequest(
	code int,
	response any,
	options *map[string]any,
) {
	defaultResponse := DefaultResponse{Data: response}
	req := w.R

	if options != nil {
		if _, ok := (*options)[key_time]; ok {
			timeValue := (*options)[key_time].(int32)
			defaultResponse.Time = &timeValue
		}

		if _, ok := (*options)[key_message]; ok {
			message := (*options)[key_message].(string)
			defaultResponse.Message = &message
		}

		if _, ok := (*options)[key_status]; ok {
			status := (*options)[key_status].(bool)
			defaultResponse.Status = &status
		}
	}

	req.WriteHeader(code)
	data, err := json.Marshal(defaultResponse)

	if err != nil {
		req.WriteHeader(500)
		_, err = req.Write([]byte(fmt.Sprint(" { \"message\": \"Server error\" }")))
		if err != nil {
			panic(err)
		}
	}

	_, err = w.R.Write(data)

	if err != nil {
		req.WriteHeader(500)
		_, err = req.Write([]byte(fmt.Sprint(" { \"message\": \"Server error\" }")))
		if err != nil {
			panic(err)
		}
	}
}
