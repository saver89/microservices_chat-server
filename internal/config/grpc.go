package config

import (
	"errors"
	"net"
	"os"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

// GRPCConfig is an interface for gRPC configuration
type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

// NewGRPCConfig creates a new gRPC configuration
func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

// Address returns the address of the gRPC server
func (g *grpcConfig) Address() string {
	return net.JoinHostPort(g.host, g.port)
}
