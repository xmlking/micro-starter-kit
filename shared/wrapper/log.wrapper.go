package wrapper

import (
	"context"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	log "github.com/sirupsen/logrus"
)

// LogWrapper is a handler wrapper to log Requests with Context metadata
func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		md, _ := metadata.FromContext(ctx)
		log.WithFields(map[string]interface{}{
			"category": "LogWrapper",
			"service":  req.Service(),
			"method":   req.Method(),
			"ctx":      md,
		}).Debug("Server-side request")
		return fn(ctx, req, rsp)
	}
}

type clientLogWrapper struct {
	client.Client
}

func (l *clientLogWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	log.WithFields(map[string]interface{}{
		"category": "LogWrapper",
		"service":  req.Service(),
		"method":   req.Method(),
		"ctx":      md,
	}).Debug("Client-side request")
	return l.Client.Call(ctx, req, rsp)
}

// NewClientLogWrapper is a client wrapper to log Requests with Context metadata
func NewClientLogWrapper(c client.Client) client.Client {
	return &clientLogWrapper{c}
}
