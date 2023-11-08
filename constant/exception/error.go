package exception

type Error struct {
	code int
	msg  string
}

func (a *Error) Error() string {
	return a.msg
}

func (a *Error) Code() int {
	return a.code
}

func NewError(code int, msg string) *Error {
	return &Error{
		code: code,
		msg:  msg,
	}
}
