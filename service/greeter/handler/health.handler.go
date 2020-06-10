package handler

import (
    "context"

    "github.com/rs/zerolog/log"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    healthPB "github.com/xmlking/micro-starter-kit/service/greeter/proto/health"
)

// Greeter struct
type healthHandler struct{}

// NewUserHandler returns an instance of `UserServiceHandler`.
func NewHealthHandler() healthPB.HealthHandler {
    return &healthHandler{}
}

func (s *healthHandler) Check(ctx context.Context, req *healthPB.HealthCheckRequest, rsp *healthPB.HealthCheckResponse) error {
    log.Info().Msg("Received Health.Check request")
    rsp.Status = healthPB.HealthCheckResponse_SERVING
    return nil
}

func (s *healthHandler) Watch(ctx context.Context, req *healthPB.HealthCheckRequest, rsp healthPB.Health_WatchStream) error {
    log.Info().Msg("Received Health.Watch request")
    return status.Errorf(codes.Unimplemented, "Unimplemented")
}
