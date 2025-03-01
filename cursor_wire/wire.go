//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"kratos_one/cursor_wire/internal/api"
	"kratos_one/cursor_wire/internal/config"
	"kratos_one/cursor_wire/internal/repository"
	"kratos_one/cursor_wire/internal/service"
)

// InitializeAPI 是一个Wire函数，用于初始化API
// 通过使用wire.Build，我们告诉Wire如何构建依赖关系图
func InitializeAPI() (*api.API, error) {
	// wire.Build会构建依赖图并生成代码来初始化API
	wire.Build(
		// 提供Config
		config.NewConfig,

		// 提供Database，需要Config
		repository.NewDatabase,

		// 提供Repository，需要Config和Database
		repository.NewRepository,

		// 提供Service，需要Repository
		service.NewService,

		// 提供API，需要Config和Service
		api.NewAPI,
	)

	// 这个返回值是虚拟的，实际的返回值将由Wire生成
	return nil, nil
}
