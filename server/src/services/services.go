package services

import (
	"github.com/enfipy/grpchat/server/src/config"
	"github.com/enfipy/grpchat/server/src/helpers"

	chatController "github.com/enfipy/grpchat/server/src/services/chat/controller"
	chatDelivery "github.com/enfipy/grpchat/server/src/services/chat/delivery"
	chatUsecase "github.com/enfipy/grpchat/server/src/services/chat/usecase"

	"github.com/enfipy/locker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func InitServices(cnfg *config.Config) (*grpc.Server, func()) {
	locker := locker.Initialize()
	pool := helpers.InitRedis(cnfg)
	srv := initServerInstance()

	chatUcs := chatUsecase.NewUsecase(cnfg, pool, locker)
	chatCnr := chatController.NewController(cnfg, chatUcs)
	chatDelivery.NewGRPC(srv, chatCnr)

	reflection.Register(srv)

	return srv, func() {
		srv.GracefulStop()
	}
}

func initServerInstance() *grpc.Server {
	uIntOpt := grpc.UnaryInterceptor(helpers.UnaryServerInterceptor)
	sIntOpt := grpc.StreamInterceptor(helpers.StreamServerInterceptor)

	return grpc.NewServer(uIntOpt, sIntOpt)
}
