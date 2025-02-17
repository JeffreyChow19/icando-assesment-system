package httperror

import (
	"github.com/pkg/errors"
	"net/http"
)

// when using this always use http.StatusBadRequest
// use this when validating fields,
// e.g. (used for validating with gin binding tag
//
//	if err != nil {
//	       var ve validator.ValidationErrors
//	       if errors.As(err, &ve) {
//	           out := make([]FieldError, len(ve))
//	           for i, fe := range ve {
//	               out[i] = FieldError{fe.Field(), msgForTag(fe.Tag())}
//	           }
//	           c.JSON(http.StatusBadRequest, gin.H{"errors": out})
//	       }
//	       return
//	   }

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Tag     string `json:"tag"`
}

func MsgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return ""
}

type HttpError struct {
	StatusCode int
	Err        error
}

func (h HttpError) Error() string {
	return h.Err.Error()
}

var InternalServerError = &HttpError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("Internal Server Error"),
}

var UnauthorizedError = &HttpError{
	StatusCode: http.StatusUnauthorized,
	Err:        errors.New("Unauthorized"),
}

var ForbiddenError = &HttpError{
	StatusCode: http.StatusForbidden,
	Err:        errors.New("Forbidden"),
}
