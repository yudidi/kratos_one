package service

import (
	"kratos_one/cursor_wire/internal/repository"
)

// Service 提供业务逻辑
type Service struct {
	repo *repository.Repository
}

// NewService 创建一个新的Service实例
// 注意Service依赖于Repository，这个依赖将被Wire注入
func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetUserInfo 获取用户信息
func (s *Service) GetUserInfo(id int) string {
	// 调用repository获取用户信息，并添加一些业务逻辑
	userInfo := s.repo.GetUserByID(id)
	return "Service processed: " + userInfo
}
