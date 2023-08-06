//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"f-dorm/app/server/internal/biz"
	"f-dorm/app/server/internal/data"
	"f-dorm/app/server/internal/server"
	"f-dorm/app/server/internal/service"

	"f-dorm/core/app"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp() (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, app.ProviderSet, newApp))
}
