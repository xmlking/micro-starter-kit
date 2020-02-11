package validator

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2/server"
	"github.com/xmlking/micro-starter-kit/shared/errors"
)

type Validator interface {
	Validate() error
}

func NewHandlerWrapper() server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			if v, ok := req.Body().(Validator); ok { // Don’t panic!
				if err := v.Validate(); err != nil {
					return errors.ValidationError(fmt.Sprintf("%s.%s", req.Service(), req.Method()), "validation error: %v", err)
				}
			}
			return fn(ctx, req, rsp)
		}
	}
}
