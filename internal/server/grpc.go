package server

import (
	v1 "f-dorm/api/demo/v1"
	"f-dorm/app/demo/internal/conf"
	"f-dorm/app/demo/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(greeter *service.DemoService) *grpc.Server {
	c := &conf.Server{}
	if err := viper.UnmarshalKey("server", &c); err != nil {
		panic(err)
	}
	grpcSecurityKey := viper.GetString("security_key")
	if grpcSecurityKey == "" {
		panic("need grpc security key")
	}
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			metadata.Server(),
			jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(grpcSecurityKey), nil
			}),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(time.Duration(*c.Grpc.Timeout)*time.Second))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterDemoServer(srv, greeter)
	return srv
}
