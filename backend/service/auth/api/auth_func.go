package api

import (
	"chat_app/service/auth/function"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

var _ function.AuthFunction = (*AuthFunction)(nil)

type AuthFunction struct {
	function.AuthFunction
	logx.Logger
	AuthService *AuthService
}

func NewAuthFunction(svc *AuthService) *AuthFunction {
	ctx := context.Background()

	return &AuthFunction{
		Logger:      logx.WithContext(ctx),
		AuthService: svc,
	}
}

func (contractFunc *AuthFunction) Start() error {
	return nil
}
