package delivery

import (
	"github.com/enfipy/grpchat/server/src/services/chat"

	chatPB "github.com/enfipy/grpchat/schema/gen/go"

	"google.golang.org/grpc"
)

func NewDelivery(serverInstance *grpc.Server, chatController chat.Controller) {
	server := &ChatServer{
		ChatController: chatController,
	}

	chatPB.RegisterChatServer(serverInstance, server)
}
