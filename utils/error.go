package utils

type HttpError struct{
	Status int
	Err string
}

func NewHttpError(status int, err string) *HttpError {
	return &HttpError{
		Status: status,
		Err: err,
	}
}