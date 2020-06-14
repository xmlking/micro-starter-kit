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

func NewSubscriberWrapper() server.SubscriberWrapper {
    return func(fn server.SubscriberFunc) server.SubscriberFunc {
        return func(ctx context.Context, p server.Message) error {
            if v, ok := p.Payload().(Validator); ok { // Don’t panic!
                if err := v.Validate(); err != nil {
                    return errors.ValidationError(fmt.Sprintf("%s", p.Topic()), "validation error: %v", err)
                }
            }
            return fn(ctx, p)
        }
    }
}

// TODO: implement client-side Validator
//type clientWrapper struct {
//   client.Client
//}
//func (l *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) (err error) {
//    return
//}
//
//func (l *clientWrapper) Publish(ctx context.Context, p client.Message, opts ...client.PublishOption) (err error) {
//    return
//}
//func NewClientWrapper() client.Wrapper {
//    return func(c client.Client) client.Client {
//        return &clientWrapper{c}
//    }
//}
