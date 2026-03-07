package model

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   error  `json:"error"`
}
