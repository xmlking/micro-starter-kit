package handler

import (
    "context"
    "log"

    purePB "github.com/xmlking/micro-starter-kit/service/pure/proto/pure"
)


type pureHandler struct {
}


func NewPureHandler() purePB.PingServer {
    return &pureHandler{}
}

func (s *pureHandler) SayHello(ctx context.Context, in *purePB.PingMessage) (*purePB.PingMessage, error) {
    log.Printf("Receive message %s", in.Greeting)
    return &purePB.PingMessage{Greeting: "bar"}, nil
}
