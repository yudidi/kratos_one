package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"kratos_one/cursor_wire/internal/config"
	"kratos_one/cursor_wire/internal/service"
)

// API 表示应用程序的API层
type API struct {
	config  *config.Config
	service *service.Service
}

// NewAPI 创建一个新的API实例
// 注意API依赖于Config和Service，这些依赖将被Wire注入
func NewAPI(cfg *config.Config, svc *service.Service) *API {
	return &API{
		config:  cfg,
		service: svc,
	}
}

// Start 启动API服务器
func (a *API) Start() error {
	addr := fmt.Sprintf("%s:%d", a.config.Server.Host, a.config.Server.Port)

	http.HandleFunc("/user/", a.handleGetUser)

	log.Printf("Starting server on %s\n", addr)
	return http.ListenAndServe(addr, nil)
}

// handleGetUser 处理获取用户信息的HTTP请求
func (a *API) handleGetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 从URL中提取用户ID
	idStr := r.URL.Path[len("/user/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// 调用服务层获取用户信息
	userInfo := a.service.GetUserInfo(id)

	// 返回用户信息
	fmt.Fprint(w, userInfo)
}
