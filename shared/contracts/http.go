package contracts

// ! This file defines your standard API response format, which all your microservices and API Gateway will use.

// APIResponse is the response structure for the API.
type APIResponse struct {
	Data  any       `json:"data,omitempty"`
	Error *APIError `json:"error,omitempty"`
}

// APIError is the error structure for the API.
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
