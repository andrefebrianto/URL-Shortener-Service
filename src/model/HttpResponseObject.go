package domain

type HttpResponseObject struct {
	Message string `json:"message"`
	Data    interface{}
}
