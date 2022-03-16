package error

type code int

const (
	internal          code = 1
	notFound          code = 2
	invalidData       code = 3
	invalidCredential code = 4
)

func newError(code code, message string, debugError error) IError {
	return &err{errorCode: code, debugError: debugError, Message: message}
}
