package models

// ErrorResponse representa uma resposta de erro na API
type ErrorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}
