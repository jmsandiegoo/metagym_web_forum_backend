package api

type ErrNotAuthenticated struct {
	Err error
}

func (e ErrNotAuthenticated) Error() string {
	return "Authentication Required"
}

func (e ErrNotAuthenticated) Unwrap() error {
	return e.Err
}

// type ErrInvalidCredentials struct {
// 	Err error
// }

// func (e ErrInvalidCredentials) Error() string {
// 	return "Invalid Credentials"
// }

// func (e ErrInvalidCredentials) Unwrap() error {
// 	return e.Err
// }

type ErrUser struct {
	Message      string
	ErrorContent map[string]interface{}
	Err          error
}

func (e ErrUser) Error() string {
	return e.Message
}

// Todo stringify map

func (e ErrUser) Unwrap() error {
	return e.Err
}
