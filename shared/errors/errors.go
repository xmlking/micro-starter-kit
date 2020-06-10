package errors

import (
	"fmt"
	"net/http"

	"github.com/micro/go-micro/v2/errors"
)

// ErrorCode for app
type ErrorCode int

// ErrorDetail for app
type ErrorDetail struct {
	ID     string
	Detail string
	Code   int32
}

const (
	// EC1 represents there is an error1
	EC1 ErrorCode = iota
	// EC2 represents there is an error2
	EC2
	// EC3 represents there is an error3
	EC3
	// EC4 represents there is an error4
	EC4
	// SME SendMailError
	SME
	// DBE DatabaseError
	DBE
	// PSE PubSubError
	PSE
)

var appErrors = map[ErrorCode]ErrorDetail{
	EC1: {"EC1", "not good", 500},
	EC2: {"EC2", "not valid", 500},
	EC3: {"EC3", "not valid", 500},
	EC4: {"EC4", "not valid", 500},
	SME: {"SME", "unable to send email: %v", 500},
	DBE: {"DBE", "database error: %v", 500},
	PSE: {"PSE", "broker publish error: %v", 500},
}

// TODO: Should I use https://github.com/grpc/grpc/blob/master/doc/http-grpc-status-mapping.md
// https://github.com/avinassh/grpc-errors/blob/master/go/server.go
// http://avi.im/grpc-errors/

// AppError - App specific Error
func AppError(errorCode ErrorCode, a ...interface{}) error {
	return &errors.Error{
		Id:     appErrors[errorCode].ID,
		Code:   appErrors[errorCode].Code,
		Detail: fmt.Sprintf(appErrors[errorCode].Detail, a...),
		Status: http.StatusText(500),
	}
}

// ValidationError - Unprocessable Entity
func ValidationError(id, format string, a ...interface{}) error {
	return &errors.Error{
		Id:     id,
		Code:   422,
		Detail: fmt.Sprintf(format, a...),
		Status: http.StatusText(422),
	}
}
