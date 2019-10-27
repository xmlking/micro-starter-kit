package transaction

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	log "github.com/sirupsen/logrus"
)

type TransationEvent struct {
	Rsp []byte
	Req []byte
}

// add TranID to context if not present
func addTranIDToContext(ctx context.Context) context.Context {
	md, _ := metadata.FromContext(ctx)
	if tranId := md["trans-id"]; len(tranId) == 0 {
		md = metadata.Copy(md)
		md["trans-id"] = uuid.New().String()
		return metadata.NewContext(ctx, md)
	} else {
		return ctx
	}
}
func publish(ctx context.Context, publisher micro.Publisher, req, rsp proto.Message) (err error) {
	reqB, err := proto.Marshal(req)
	if err != nil {
		log.WithError(err).Errorf("marshaling error for req")
		return
	}
	resB, err := proto.Marshal(rsp)
	if err != nil {
		log.WithError(err).Errorf("marshaling error for rsp")
		return
	}
	event := &TransationEvent{}
	event.Req = reqB
	event.Rsp = resB
	if err = publisher.Publish(ctx, event); err != nil {
		log.WithError(err).Errorf("Publisher: Failed publishing transation")
	}
	return
}

// NewHandlerWrapper return Log HandlerWrapper which publish transaction event
func NewHandlerWrapper(p micro.Publisher) server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) (err error) {
			ctx = addTranIDToContext(ctx)
			err = fn(ctx, req, rsp)
			// we already logged error in Publish. lets ignore error here.
			_ = publish(ctx, p, req.Body().(proto.Message), rsp.(proto.Message))
			// go publish(ctx, p, req.Body().(proto.Message), rsp.(proto.Message))
			return
		}
	}
}
