package errors

type ApiError struct {
	Err    error
	Msg    string
	Status int
}

func NewApiError(err error, msg string, status int) ApiError {
	return ApiError{
		Err:    err,
		Msg:    msg,
		Status: status,
	}
}

func (e ApiError) Error() string {
	return e.Err.Error()
}

type UINofifyError struct {
	Msg    string
	Status int
}

func NewUINotifyError(msg string, status int) UINofifyError {
	return UINofifyError{
		Msg:    msg,
		Status: status,
	}
}

func (e UINofifyError) Error() string {
	return e.Msg
}
