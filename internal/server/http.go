package server

import (
	v1 "f-dorm/api/demo/v1"
	"f-dorm/app/demo/internal/conf"
	"f-dorm/app/demo/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/spf13/viper"
	"time"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(greeter *service.DemoService, logger log.Logger) *http.Server {
	c := &conf.Server{}
	if err := viper.UnmarshalKey("server", &c); err != nil {
		panic(err)
	}
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(time.Duration(*c.Http.Timeout)*time.Second))
	}
	srv := http.NewServer(opts...)
	v1.RegisterDemoHTTPServer(srv, greeter)
	return srv
}
