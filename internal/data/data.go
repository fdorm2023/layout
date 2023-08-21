package data

import (
	v1 "f-dorm/api/user/v1"
	"f-dorm/core/app"
	"f-dorm/core/const/codename"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	UserClient v1.UserClient
}

// NewData .
func NewData(logger log.Logger, r registry.Discovery) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	mapClient := app.InitGRPCClient(r, []string{}, codename.ServiceUser)
	userClient := v1.NewUserClient(mapClient[codename.ServiceUser])
	return &Data{
		UserClient: userClient,
	}, cleanup, nil
}
