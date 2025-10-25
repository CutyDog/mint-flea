package errors

import (
	"encoding/json"
	"net/http"
)

// GraphQLError represents a GraphQL error response
type GraphQLError struct {
	Message    string                 `json:"message"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

// GraphQLResponse represents a GraphQL response with errors
type GraphQLResponse struct {
	Data   interface{}    `json:"data,omitempty"`
	Errors []GraphQLError `json:"errors,omitempty"`
}

// SendGraphQLError sends a GraphQL-compatible error response
func SendGraphQLError(w http.ResponseWriter, message string, code string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := GraphQLResponse{
		Data: nil, // エラー時はdataをnullにする
		Errors: []GraphQLError{
			{
				Message: message,
				Extensions: map[string]interface{}{
					"code": code,
				},
			},
		},
	}

	json.NewEncoder(w).Encode(response)
}

// SendUnauthenticatedError sends an authentication error
func SendUnauthenticatedError(w http.ResponseWriter, message string) {
	SendGraphQLError(w, message, UNAUTHENTICATED, http.StatusUnauthorized)
}

// SendForbiddenError sends a forbidden error
func SendForbiddenError(w http.ResponseWriter, message string) {
	SendGraphQLError(w, message, FORBIDDEN, http.StatusForbidden)
}

// SendBadRequestError sends a bad request error
func SendBadRequestError(w http.ResponseWriter, message string) {
	SendGraphQLError(w, message, BAD_REQUEST, http.StatusBadRequest)
}

// SendInternalError sends an internal server error
func SendInternalError(w http.ResponseWriter, message string) {
	SendGraphQLError(w, message, INTERNAL_ERROR, http.StatusInternalServerError)
}
