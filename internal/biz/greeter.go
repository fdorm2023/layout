package biz

import (
	"context"
	user_v1 "f-dorm/api/user/v1"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

type GreeterRepo interface {
	GetUserByUserName(ctx context.Context, userName string) (*user_v1.GetProfileByUserNameResponse, error)
}

type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GreeterUsecase) GenerateHelloMessage(ctx context.Context, userName string) (string, error) {
	uc.log.WithContext(ctx).Infof("Generate hello message: %s", userName)
	user, err := uc.repo.GetUserByUserName(ctx, userName)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Xin chào %s %s", user.FirstName, user.LastName), nil
}
