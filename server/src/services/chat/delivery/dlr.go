package delivery

import (
	"github.com/enfipy/grpchat/server/src/services/chat"

	chatPB "github.com/enfipy/grpchat/schema/gen/go"

	"google.golang.org/grpc"
)

type ChatServer struct {
	ChatController chat.Controller
}

func NewDelivery(serverInstance *grpc.Server, chatController chat.Controller) {
	server := &ChatServer{
		ChatController: chatController,
	}

	chatPB.RegisterChatServer(serverInstance, server)
}
