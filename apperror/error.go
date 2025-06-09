// file error.go ini digunnakan untuk handle error dari input dari user

package apperror

import (
	"errors"
	"fmt"
)

// Error types
var (
	ErrBadRequest = errors.New("bad request")
	ErrNotFound   = errors.New("not found")
	ErrInternal   = errors.New("internal server error")
	ErrConflict   = errors.New("conflict")
)

// CustomError adalah struct untuk error kustom
type CustomError struct {
	Code    int
	Message string
}

// Function untuk membuat error baru dari error kustom
func (e *CustomError) Error() string {
	return e.Message
}

// ValidationError membuat error baru untuk validasi
func ValidationError(field string) error {
	return &CustomError{ // Menggunakan error kustom agar kesal error input user code 400
		Code:    400,
		Message: fmt.Sprintf("%s'", field),
	}
}

func ErrorWithMessage(message string) error {
	return &CustomError{ // Menggunakan error kustom agar kesal error input user code 400
		Code:    400,
		Message: fmt.Sprintf("%s'", message),
	}
}

func InternalServerErrorWithMessage(message string) error {
	return &CustomError{
		Code:    500,
		Message: message,
	}
}

func ConflictError(message string) error {
	return &CustomError{
		Code:    409,
		Message: message,
	}
}

// DetermineErrorType mengecek jenis error dan balikin status + message
func DetermineErrorType(err error) (int, string) {
	switch e := err.(type) {
	case *CustomError:
		return e.Code, e.Message
	default:
		switch {
		case errors.Is(err, ErrBadRequest):
			return 400, "Bad Request"
		case errors.Is(err, ErrNotFound):
			return 404, "Not Found"
		case errors.Is(err, ErrInternal):
			return 409, "Conflict"
		default:
			return 500, "Internal Server Error"
		}
	}
}
