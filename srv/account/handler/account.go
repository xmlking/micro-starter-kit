package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	account "github.com/xmlking/micro-starter-kit/srv/account/proto/account"
)

type Account struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Account) Call(ctx context.Context, req *account.Request, rsp *account.Response) error {
	log.Log("Received Account.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Account) Stream(ctx context.Context, req *account.StreamingRequest, stream account.Account_StreamStream) error {
	log.Logf("Received Account.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&account.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Account) PingPong(ctx context.Context, stream account.Account_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&account.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
