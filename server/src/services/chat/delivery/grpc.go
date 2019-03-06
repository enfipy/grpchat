package delivery

import (
	"github.com/enfipy/grpchat/server/src/services/chat"

	chatPB "github.com/enfipy/grpchat/schema/gen/go"

	"golang.org/x/net/context"
)

type ChatServer struct {
	ChatController chat.Controller
}

func (server *ChatServer) GetMessages(req *chatPB.GetMessagesRequest, stream chatPB.Chat_GetMessagesServer) error {
	ctx := stream.Context()

	userLeft := server.ChatController.UserJoin(req.Username)
	defer userLeft()

	done := make(chan struct{})
	go func() {
		defer close(done)

		outgoing := server.ChatController.GetMessages()
		for {
			select {
			case <-ctx.Done():
				return
			case out, ok := <-outgoing:
				if !ok {
					return
				}
				stream.Send(out.ToProtobuf())
			}
		}
	}()

	<-done
	return nil
}

func (server *ChatServer) SendMessage(ctx context.Context, req *chatPB.ClientMessage) (*chatPB.SendMessageResponse, error) {
	// Todo: Authenticate user
	username := getUsernameFromContext(ctx)
	server.ChatController.SendMessage(username, req.Message)
	return new(chatPB.SendMessageResponse), nil
}
