# Go Wire 依赖注入教程

这个项目是一个关于 Google Wire 依赖注入工具的教程示例。通过本项目，你将学习到：

1. 什么是依赖注入以及为什么需要它
2. 如何使用 Wire 来简化依赖注入
3. Wire 的实际应用场景和最佳实践

## 项目结构

```
.
├── README.md              # 项目说明
├── go.mod                 # Go 模块定义
├── go.sum                 # 依赖版本锁定文件
├── main.go                # 主程序入口
├── wire.go                # Wire 依赖注入定义
├── wire_gen.go            # Wire 自动生成的代码
└── internal/              # 内部包
    ├── config/            # 配置相关
    ├── repository/        # 数据库访问层
    ├── service/           # 业务逻辑层
    └── api/               # API 层
```

## 如何使用

1. 安装 Wire 工具：`go install github.com/google/wire/cmd/wire@latest`
2. 查看项目代码和注释，了解 Wire 的使用方法
3. 运行示例：`go run main.go`
4. 修改依赖关系后重新生成 Wire 代码：`wire`

## 依赖

- Go 1.18 或更高版本
- github.com/google/wire 