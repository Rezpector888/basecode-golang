package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/common/config"
	"example.com/common/middlewares"
	"example.com/models/seeders"
	"example.com/routes"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	gin.SetMode(config.AppConfig.AppMode)

	router := gin.Default()

	config.ConnectDatabase()

	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(middlewares.SecureMiddleware())

	//seeding() // Uncoment this for seeding
	routes.SetupRoute(router)

	addr := fmt.Sprintf(":%s", config.AppConfig.Port)
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	config.LogMessage("INFO", "Server is running on port "+config.AppConfig.Port)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			config.LogMessage("ERROR", "Server error: "+err.Error())
		}
	}()

	waitForShutdown(server)
}

func waitForShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	config.LogMessage("INFO", "Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		config.LogMessage("ERROR", "Server forced to shutdown: "+err.Error())
	} else {
		config.LogMessage("INFO", "Server stopped successfully")
	}
}

func seeding() {
	seeders.SeedUser()
}
