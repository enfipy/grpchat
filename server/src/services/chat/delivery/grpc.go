package delivery

import (
	"github.com/enfipy/grpchat/server/src/helpers"
	"github.com/enfipy/grpchat/server/src/services/chat"

	chatPB "github.com/enfipy/grpchat/schema/gen/go"

	"google.golang.org/grpc"
)

type ChatServer struct {
	ChatController chat.Controller
}

func NewGRPC(serverInstance *grpc.Server, chatController chat.Controller) {
	server := &ChatServer{
		ChatController: chatController,
	}

	chatPB.RegisterChatServer(serverInstance, server)
}

func (server *ChatServer) MessageStream(stream chatPB.Chat_MessageStreamServer) error {
	username := getUsernameFromContext(stream.Context())

	userLeft := server.ChatController.UserJoin(username)
	defer userLeft()

	done := make(chan struct{})
	go func() {
		defer close(done)

		incoming := make(chan string)
		defer close(incoming)

		streamChannel := helpers.ToChan(stream, func() interface{} { return new(chatPB.ClientMessage) })
		outgoing := server.ChatController.MessageStream(username, incoming)

		for {
			select {
			case out, ok := <-outgoing:
				if !ok {
					return
				}
				stream.Send(out.ToProtobuf())
			case in, ok := <-streamChannel:
				if !ok {
					return
				}
				msg := in.(*chatPB.ClientMessage)
				incoming <- msg.Message
			}
		}
	}()

	<-done
	return nil
}
