package memos

type NotFoundError struct {
	Msg string
}

func (e *NotFoundError) Error() string {
	return e.Msg
}

type AlreadyExistsError struct {
	Msg string
}

func (e *AlreadyExistsError) Error() string {
	return e.Msg
}

type IllegalDeviceTokenError struct {
	Msg string
}

func (e *IllegalDeviceTokenError) Error() string {
	return e.Msg
}
