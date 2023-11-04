package dto

type Base struct {
	Success bool   `json:"success" default:"false"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    any    `json:"data"`
}
