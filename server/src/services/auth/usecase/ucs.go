package usecase

import (
	"github.com/enfipy/grpchat/server/src/config"
	"github.com/enfipy/grpchat/server/src/services/auth"

	"github.com/enfipy/locker"
	"github.com/gomodule/redigo/redis"
)

type authUsecase struct {
	config *config.Config
	pool   *redis.Pool
	locker *locker.Locker
}

func NewUsecase(cnfg *config.Config, pool *redis.Pool, locker *locker.Locker) auth.Usecase {
	return &authUsecase{
		config: cnfg,
		pool:   pool,
		locker: locker,
	}
}
