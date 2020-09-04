package handler

import (
	"encoding/json"
	"net/http"
)

const (
	// Error .
	Error = "error"
	// Message .
	Message = "message"
)

// Response .
type Response struct {
	MessageType string      `json:"message-type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

// NewResponse .
func NewResponse(msjT, msj string, data interface{}) *Response {
	return &Response{msjT, msj, data}
}

// JSON .
func (r *Response) JSON(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	data := Response{r.MessageType, r.Message, r.Data}
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
