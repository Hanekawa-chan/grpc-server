package grpc_server

import "time"

import (
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// NewGRPCServer returns grpc server by config with middlewares
func NewGRPCServer(cfg *Config, middlewares ...grpc.UnaryServerInterceptor) *grpc.Server {
	interceptors := []grpc.UnaryServerInterceptor{
		grpcRecovery.UnaryServerInterceptor(grpcRecovery.WithRecoveryHandlerContext(recoveryHandler)),
		logIncomingRequestsMiddleware,
		grpcPrometheus.UnaryServerInterceptor,
		grpcValidator.UnaryServerInterceptor(),
	}

	interceptors = append(interceptors, middlewares...)
	return grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: cfg.MaxConnectionIdle * time.Minute,
		Timeout:           cfg.Timeout * time.Second,
		MaxConnectionAge:  cfg.MaxConnectionAge * time.Minute,
		Time:              cfg.Timeout * time.Minute,
	}), grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(interceptors...)))
}
