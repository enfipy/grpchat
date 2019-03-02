package main

import (
	"log"
	"os"

	"github.com/enfipy/grpchat/server/test/chat"

	"google.golang.org/grpc"
)

func main() {
	log.Print("Test started")

	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatal("Valid service name and rpc method must be provided")
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	checkService(args, conn)
}

func checkService(args []string, conn *grpc.ClientConn) {
	switch args[0] {
	case "chat":
		chat.ExecMethod(args[1], conn)
	default:
		os.Exit(1)
	}
}
