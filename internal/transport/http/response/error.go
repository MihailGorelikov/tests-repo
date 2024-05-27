package response

// Error is an error response.
type Error struct {
	// StatusCode is the status code.
	StatusCode int `json:"statusCode"`
	// Message is the error message.
	Message string `json:"message"`
}
