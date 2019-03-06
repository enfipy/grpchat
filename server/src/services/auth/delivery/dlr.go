package delivery

import (
	"log"

	"github.com/enfipy/grpchat/server/src/services/auth"

	"google.golang.org/grpc"
)

type AuthServer struct {
	AuthController auth.Controller
}

func NewDelivery(serverInstance *grpc.Server, authController auth.Controller) {
	server := &AuthServer{
		AuthController: authController,
	}

	// Todo: Implement auth
	log.Print(server)
}
