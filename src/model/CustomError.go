package domain

type CustomError struct {
	Message string `json:"message"`
	Data    interface{}
}
