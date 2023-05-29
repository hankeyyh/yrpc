package server

import "context"

// Server 通用的server，可以代表一个rpc server，也可以代表一个http server
type Server interface {
	Listen(ctx context.Context) error
	Serve(ctx context.Context) error
	Stop(ctx context.Context) error
	GracefulStop(ctx context.Context) error
}
