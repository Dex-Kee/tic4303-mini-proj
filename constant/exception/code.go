package exception

var (
	ErrLoginFailed   = NewError(10000, "login failed, username or password is incorrect")
	ErrLockout       = NewError(10001, "account is locked, please try again later in 600 seconds")
	ErrUpdateInvalid = NewError(10100, "invalid update operation")
)
