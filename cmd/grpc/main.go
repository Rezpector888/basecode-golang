package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"example.com/common/config"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	config.LoadConfig()

	gin.SetMode(config.AppConfig.AppMode)

	config.ConnectDatabase()

	addr := fmt.Sprintf(":%s", config.AppConfig.GrpcPort)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		config.LogMessage("ERROR", "GRPC error running: "+err.Error())
	}
	serverGrpc := grpc.NewServer()

	config.LogMessage("INFO", "GRPC is running on port "+config.AppConfig.GrpcPort)

	if err := serverGrpc.Serve(listen); err != nil {
		config.LogMessage("ERROR", "GRPC error running: "+err.Error())
	}

	waitForShutdown(serverGrpc)
}

func waitForShutdown(server *grpc.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	config.LogMessage("INFO", "Shutting down server...")

	server.Stop()

	config.LogMessage("INFO", "Server stopped successfully")
}
