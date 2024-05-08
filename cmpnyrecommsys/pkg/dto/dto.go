package dto

import (
)

//MessagResponse
type MessageResponse struct {
	Message string `json:"message"`
}

//ErrorResponse
type ErrorResponse struct {
	Error string `json:"error"`
}