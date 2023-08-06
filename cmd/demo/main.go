package main

import (
	"f-dorm/core/const/codename"
	"github.com/go-kratos/kratos/v2/registry"
	"net/url"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	fdorm_app "f-dorm/core/app"
	_ "go.uber.org/automaxprocs"
)

func newApp(app *fdorm_app.App, logger log.Logger, hs *http.Server, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(app.Id),
		kratos.Name(codename.ServiceDemo),
		kratos.Version(app.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
		kratos.Registrar(rr),
		kratos.Endpoint(&url.URL{
			Host:   app.GrpcEndpoint,
			Scheme: app.GrpcSchema,
		}),
	)
}

func main() {

	app, cleanup, err := wireApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
