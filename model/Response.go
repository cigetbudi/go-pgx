package model

type Response struct {
	StatusCode  string      `json:"status_code"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}
