package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/enfipy/grpchat/server/src/config"
	"github.com/enfipy/grpchat/server/src/services"
)

func main() {
	cnfg := config.InitConfig()

	if cnfg.AppEnv != "production" {
		log.SetFlags(0)
	} else {
		log.Print("Service started")
	}

	port := ":" + cnfg.ServerPort
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	srv, close := services.InitServices(cnfg)
	go gracefulShutdown(close)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func gracefulShutdown(close func()) {
	quitChan := make(chan os.Signal, 1)

	signal.Notify(quitChan, syscall.SIGTERM)
	signal.Notify(quitChan, syscall.SIGINT)
	signal.Notify(quitChan, syscall.SIGKILL)

	<-quitChan
	close()
}
