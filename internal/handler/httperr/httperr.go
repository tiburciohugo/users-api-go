package httperr

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error,omitempty"`
	Code    int      `json:"code"`
	Fields  []Fields `json:"fields,omitempty"`
}

type Fields struct {
	Field   string      `json:"field"`
	Value   interface{} `json:"value,omitempty"`
	Message string      `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(m, e string, c int, f []Fields) *RestErr {
	return &RestErr{
		Message: m,
		Err:     e,
		Code:    c,
		Fields:  f,
	}
}

func NewBadRequestError(m string) *RestErr {
	return &RestErr{
		Message: m,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(m string) *RestErr {
	return &RestErr{
		Message: m,
		Err:     "unauthorized_request",
		Code:    http.StatusUnauthorized,
	}
}

func NewBadRequestValidationError(m string, f []Fields) *RestErr {
	return &RestErr{
		Message: m,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Fields:  f,
	}
}

func NewInternalServerError(m string) *RestErr {
	return &RestErr{
		Message: m,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(m string) *RestErr {
	return &RestErr{
		Message: m,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(m string) *RestErr {
	return &RestErr{
		Message: m,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
