package grpc_server

import "time"

// Config contains connection settings for grpc server
type Config struct {
	MaxConnectionIdle time.Duration // MaxConnectionIdle counts in minutes
	Timeout           time.Duration // Timeout counts in seconds
	MaxConnectionAge  time.Duration // MaxConnectionAge counts in minutes
}
