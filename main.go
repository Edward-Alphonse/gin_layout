package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"gin_layout/internal/biz/router"
	"gin_layout/internal/config"
	"gin_layout/internal/db"
	"gin_layout/internal/redis"
)

var (
	configPath = flag.String("config", "./config/config.yaml", "User Account Service address")
)

func main() {
	flag.Parse()
	cfg := config.Init(*configPath)
	db.Init(cfg.Mysql)
	redis.Init(cfg.Redis)

	port := cfg.App.Port
	r := gin.Default()
	r.GET("/app1/hello", router.Hello)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil { // listen and serve on 0.0.0.0:8080
			log.Printf("listen: %s\n", err)
		}
	}()

	gracefulStop(srv)
}

// 优雅关闭
func gracefulStop(srv *http.Server) {
	// 监听终止信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// 设置优雅关闭超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 关闭服务并等待请求处理完成
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exited gracefully")
}
