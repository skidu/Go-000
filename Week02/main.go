package main

import (
	"Week02/database"
	"Week02/router"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	database.InitMySQL()

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)

	server := &http.Server{
		Addr:    `127.0.0.1:996`,
		Handler: getApp(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("服务启动失败: %v\n", err)
		}
	}()

	<-stopChan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("强制退出服务: %v\n", err)
	}

	log.Println("服务已退出")
}

func getApp() *gin.Engine {
	app := gin.Default()
	app.GET(`/user/:id`, router.Detail)
	return app
}
