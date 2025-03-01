package config

// Config 表示应用程序的配置
type Config struct {
	// 数据库配置
	Database struct {
		Driver   string
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}

	// 服务器配置
	Server struct {
		Host string
		Port int
	}
}

// NewConfig 创建一个新的配置实例
func NewConfig() *Config {
	// 在实际应用中，这里可能会从文件或环境变量中读取配置
	// 为了简单起见，我们直接返回硬编码的配置
	cfg := &Config{}

	// 设置数据库配置
	cfg.Database.Driver = "mysql"
	cfg.Database.Host = "localhost"
	cfg.Database.Port = 3306
	cfg.Database.User = "root"
	cfg.Database.Password = "password"
	cfg.Database.DBName = "example"

	// 设置服务器配置
	cfg.Server.Host = "localhost"
	cfg.Server.Port = 8080

	return cfg
}
