package biz

import (
	"context"
	user_v1 "f-dorm/api/user/v1"
	"f-dorm/core/logger"
	"fmt"
)

type GreeterRepo interface {
	GetUserByUserName(ctx context.Context, userName string) (*user_v1.GetProfileByUserNameResponse, error)
}

type GreeterUsecase struct {
	repo GreeterRepo
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo) *GreeterUsecase {
	return &GreeterUsecase{repo: repo}
}

func (uc *GreeterUsecase) GenerateHelloMessage(ctx context.Context, userName string) (string, error) {
	ctxLogger := logger.NewLogger(ctx)
	ctxLogger.Infof("Generate hello message: %s", userName)
	user, err := uc.repo.GetUserByUserName(ctx, userName)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Xin chào %s %s", user.FirstName, user.LastName), nil
}
