package helpers

import "google.golang.org/grpc"

func ToChan(stream grpc.ServerStream, factory func() interface{}) <-chan interface{} {
	channel := make(chan interface{})

	go func() {
		defer close(channel)
		for {
			msg := factory()
			if err := stream.RecvMsg(msg); err != nil {
				return
			}
			channel <- msg
		}
	}()

	return channel
}
