package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	grpc_health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type (
	health struct{}
)

func NewHealth() grpc_health.HealthServer {
	return &health{}
}

func (h *health) Check(context.Context, *grpc_health.HealthCheckRequest) (*grpc_health.HealthCheckResponse, error) {
	return &grpc_health.HealthCheckResponse{
		Status: grpc_health.HealthCheckResponse_SERVING,
	}, nil
}

// Watch : use streming rpc
func (h *health) Watch(*grpc_health.HealthCheckRequest, grpc_health.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "not impl")
}
