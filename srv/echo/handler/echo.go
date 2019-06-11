package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	echo "github.com/xmlking/micro-starter-kit/srv/echo/proto/echo"
)

type Echo struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Echo) Call(ctx context.Context, req *echo.Request, rsp *echo.Response) error {
	log.Log("Received Echo.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Echo) Stream(ctx context.Context, req *echo.StreamingRequest, stream echo.Echo_StreamStream) error {
	log.Logf("Received Echo.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&echo.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Echo) PingPong(ctx context.Context, stream echo.Echo_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&echo.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
