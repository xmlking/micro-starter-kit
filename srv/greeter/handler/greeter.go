package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	greeterPB "github.com/xmlking/micro-starter-kit/srv/greeter/proto/greeter"
)

// Greeter struct
type Greeter struct{}

// Hello method
func (s *Greeter) Hello(ctx context.Context, req *greeterPB.Request, rsp *greeterPB.Response) error {
	log.Info("Received Greeter.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
