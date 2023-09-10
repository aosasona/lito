package errors

type APIError struct {
	Message    string
	StatusCode int
}

func NewAPIError(message string, statusCode int) *APIError {
	return &APIError{
		Message:    message,
		StatusCode: statusCode,
	}
}

func (e *APIError) Error() string {
	return e.Message
}

func (e *APIError) Code() int {
	if e.StatusCode == 0 {
		return 500
	}

	return e.StatusCode
}
