package controller

import (
	"github.com/enfipy/grpchat/server/src/config"
	"github.com/enfipy/grpchat/server/src/services/auth"
)

type chatController struct {
	config      *config.Config
	authUsecase auth.Usecase
}

func NewController(cnfg *config.Config, authUsecase auth.Usecase) auth.Controller {
	return &chatController{
		config:      cnfg,
		authUsecase: authUsecase,
	}
}
