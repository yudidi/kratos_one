package main

import (
	"log"
)

func main() {
	// 使用Wire生成的InitializeAPI函数初始化API
	// 所有的依赖注入都会自动处理
	api, err := InitializeAPI()
	if err != nil {
		log.Fatalf("Failed to initialize API: %v", err)
	}

	// 启动API服务器
	if err := api.Start(); err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}
}
