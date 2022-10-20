package errors

type AppError struct {
	Service string `json:"service"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func New(service, code, message string) *AppError {
	return &AppError{service, code, message}
}
