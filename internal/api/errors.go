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

type ErrNotAuthorized struct {
	Err error
}

func (e ErrNotAuthorized) Error() string {
	return "No permission to carry out the request"
}

func (e ErrNotAuthorized) Unwrap() error {
	return e.Err
}

type ErrUser struct {
	Message      string
	ErrorContent map[string]interface{}
	Err          error
}

func (e ErrUser) Error() string {
	var res string
	res = e.Message
	if e.Err != nil {
		res += " " + e.Err.Error()
	}
	return res
}

// Todo stringify map

func (e ErrUser) Unwrap() error {
	return e.Err
}
