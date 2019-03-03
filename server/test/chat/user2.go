package chat

import (
	"context"
	"io"
	"log"

	chatPB "github.com/enfipy/grpchat/schema/gen/go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (t *test) User2(conn *grpc.ClientConn) {
	client := chatPB.NewChatClient(conn)

	req := chatPB.GetMessagesRequest{
		Username: "doe",
	}
	stream, err := client.GetMessages(context.Background(), &req)
	defer stream.CloseSend()
	if err != nil {
		panic(err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a message : %v", err)
			}
			log.Printf("Message: %+v", message)
		}
	}()

	<-done
}

func (t *test) User2Send(conn *grpc.ClientConn) {
	client := chatPB.NewChatClient(conn)

	headers := metadata.Pairs(
		"username", "doe",
	)
	ctx := metadata.NewOutgoingContext(context.Background(), headers)

	req := chatPB.ClientMessage{
		Message: "message",
	}
	_, err := client.SendMessage(ctx, &req)
	if err != nil {
		panic(err)
	}
}
