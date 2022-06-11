package handler

import "encoding/json"

type errorResponse struct {
	Code    uint16   `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

type successResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func toJSON(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}
