package data

import (
	"context"
	v1 "f-dorm/api/user/v1"
	"f-dorm/app/demo/internal/biz"
)

type greeterRepo struct {
	data *Data
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
	}
}

func (g *greeterRepo) GetUserByUserName(ctx context.Context, userName string) (*v1.GetProfileByUserNameResponse, error) {
	return g.data.UserClient.GetProfileByUserName(ctx, &v1.GetProfileByUserNameRequest{Username: userName})
}
