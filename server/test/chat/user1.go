package chat

import (
	"context"
	"io"
	"log"
	"time"

	chatPB "github.com/enfipy/grpchat/schema/gen/go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (t *test) User1(conn *grpc.ClientConn) {
	client := chatPB.NewChatClient(conn)

	headers := metadata.Pairs(
		"username", "johndoe",
	)
	ctx := metadata.NewOutgoingContext(context.Background(), headers)

	stream, err := client.MessageStream(ctx)
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

	messages := []*chatPB.ClientMessage{
		{Message: "John: First message"},
		{Message: "John: Second message"},
		{Message: "John: Third message"},
	}

	for _, message := range messages {
		err := stream.Send(message)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to send a message: %v", err)
		}
		time.Sleep(1 * time.Second)
	}

	<-done
}
