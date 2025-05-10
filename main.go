package main

import (
	"github-pic/internal/infrastructure/config"
	"github-pic/internal/interfaces/http"
	"log"
)

func main() {
	// 初始化配置
	if err := config.InitConfig(); err != nil {
		log.Fatalf("配置初始化失败: %v", err)
	}

	// 启动HTTP服务器
	server := http.NewServer()
	if err := server.Run(); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
