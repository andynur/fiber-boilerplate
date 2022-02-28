package serializer

type Response struct {
	Status   int         `json:"status,omitempty"`
	Code     string      `json:"code,omitempty"`
	Message  string      `json:"message,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Paginate interface{} `json:"paginate,omitempty"`
}

type ResponseSuccess struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
