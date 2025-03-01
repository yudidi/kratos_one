package repository

import (
	"fmt"

	"kratos_one/cursor_wire/internal/config"
)

// Repository 提供数据访问功能
type Repository struct {
	config *config.Config
	db     *Database
}

// Database 表示数据库连接
type Database struct {
	DSN string
}

// NewDatabase 创建一个新的数据库连接
func NewDatabase(cfg *config.Config) (*Database, error) {
	// 在实际应用中，这里会创建真正的数据库连接
	// 为了演示，我们只是生成一个连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)

	return &Database{
		DSN: dsn,
	}, nil
}

// NewRepository 创建一个新的Repository实例
// 注意这个函数依赖于Config和Database，这些依赖将被Wire注入
func NewRepository(cfg *config.Config, db *Database) *Repository {
	return &Repository{
		config: cfg,
		db:     db,
	}
}

// GetUserByID 模拟从数据库获取用户信息
func (r *Repository) GetUserByID(id int) string {
	// 模拟数据库查询
	return fmt.Sprintf("User %d from database: %s", id, r.db.DSN)
}
