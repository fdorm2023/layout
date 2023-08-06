//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"f-dorm/app/demo/internal/biz"
	"f-dorm/app/demo/internal/data"
	"f-dorm/app/demo/internal/server"
	"f-dorm/app/demo/internal/service"

	"f-dorm/core/app"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp() (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, app.ProviderSet, newApp))
}
