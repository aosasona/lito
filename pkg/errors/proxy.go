package errors

type ProxyError struct {
	message string
	code    int
}

var (
	ErrServiceNotFound      = NewProxyError(404, "No configuration found for this service")
	ErrServiceMisconfigured = NewProxyError(500, "Service is misconfigured")
)

func NewProxyError(code int, msg string) *ProxyError {
	return &ProxyError{
		message: msg,
		code:    code,
	}
}

func (e *ProxyError) Error() string { return e.message }
func (e *ProxyError) Code() int     { return e.code }
