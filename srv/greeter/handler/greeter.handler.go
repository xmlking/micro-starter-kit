package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	greeterPB "github.com/xmlking/micro-starter-kit/srv/greeter/proto/greeter"
)

// Greeter struct
type greeterHandler struct{}

// NewUserHandler returns an instance of `UserServiceHandler`.
func NewGreeterHandler() greeterPB.GreeterHandler {
	return &greeterHandler{}
}

// Hello method
func (s *greeterHandler) Hello(ctx context.Context, req *greeterPB.Request, rsp *greeterPB.Response) error {
	log.Info("Received Greeter.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
