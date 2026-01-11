package model

type ErrorGeneral struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
