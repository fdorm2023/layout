package data

import (
	"context"
	v1 "f-dorm/api/user/v1"
	"f-dorm/app/server/internal/biz"
	"f-dorm/core/app"
	"f-dorm/core/const/codename"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
)

type greeterRepo struct {
	data       *Data
	log        *log.Helper
	userClient v1.UserClient
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger, r registry.Discovery) biz.GreeterRepo {
	mapClient := app.InitGRPCClient(r, []string{}, codename.ServiceUser)
	userClient := v1.NewUserClient(mapClient[codename.ServiceUser])
	return &greeterRepo{
		data:       data,
		log:        log.NewHelper(logger),
		userClient: userClient,
	}
}

func (g *greeterRepo) GetUserByUserName(ctx context.Context, userName string) (*v1.GetProfileByUserNameResponse, error) {
	return g.userClient.GetProfileByUserName(ctx, &v1.GetProfileByUserNameRequest{Username: userName})
}
