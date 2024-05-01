package netapi

type DefaultError struct {
	Headers    map[string]string
	StatusCode int
	Data       DefaultResponse
}

type DefaultResponse struct {
	Message *string `json:"message,omitempty"`
	Data    any     `json:"data"`
	Time    *int32  `json:"time,omitempty"`
	Status  *bool   `json:"status,omitempty"`
}
