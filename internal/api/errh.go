package api

type ErrorResponse struct {
	Error string `json:"error"`
}

func ErrHandler(_ *Context, err error) interface{} {
	return &ErrorResponse{
		Error: err.Error(),
	}
}
