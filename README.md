# 文件入口介绍
my-go-project/
├── cmd/               # 主应用程序目录
│   └── appname/       # 每个子目录对应一个可执行文件
│       └── main.go    # 主程序入口
├── internal/          # 私有代码（仅本项目可访问）
│   ├── pkg1/         # 内部包1
│   └── pkg2/         # 内部包2
├── pkg/              # 公共库代码（可被外部项目导入）
│   ├── lib1/         # 公共包1
│   └── lib2/         # 公共包2
├── api/              # API定义文件（protobuf/gRPC等）
├── web/              # 前端资源（可选）
├── configs/          # 配置文件模板
├── deployments/      # 部署脚本
├── scripts/          # 辅助脚本
├── build/            # 构建输出（通常.gitignore忽略）
├── vendor/           # 依赖副本（go mod vendor生成）
├── go.mod            # 模块定义文件
├── go.sum            # 依赖校验文件
├── Makefile          # 常用命令封装（可选）
└── README.md         # 项目说明文档