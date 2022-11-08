package dto

type SuccessResult struct {
	Code string      `json:"status"`
	Data interface{} `json:"data"`
}

type ErrorResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}