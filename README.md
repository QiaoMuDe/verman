<div align="center">

# verman 🚀

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.18-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/gitee.com/MM-Q/verman)](https://goreportcard.com/report/gitee.com/MM-Q/verman)
[![Release](https://img.shields.io/badge/release-v1.0.0-brightgreen)](https://gitee.com/MM-Q/verman/releases)

</div>

> 🎯 **轻量级 Go 版本信息管理库** - 专为简化应用程序版本管理而设计的现代化解决方案

verman 是一个功能强大且易于使用的 Go 语言版本信息管理库，支持编译时注入版本信息，提供多种格式的版本输出，帮助开发者轻松管理和展示应用版本信息。

---

## ✨ 核心特性

- 🔧 **编译时注入** - 支持通过 `-ldflags` 在编译时注入版本信息
- 📊 **多种输出格式** - 提供 5 种不同的版本信息输出格式
- 🌐 **运行时信息** - 自动获取 Go 版本和平台信息
- 🎨 **简洁 API** - 直接调用全局函数，无需构造函数
- 🚀 **零依赖** - 仅使用 Go 标准库，轻量级设计
- 🔒 **数据安全** - 防止运行时意外修改版本信息
- ⚡ **高性能** - 优化的字符串格式化，支持并发访问

## 📦 安装指南

### 从 Gitee 仓库安装

```bash
go get gitee.com/MM-Q/verman
```

### 验证安装

```bash
go mod tidy
```

## 🚀 使用示例

### 基础用法

```go
package main

import (
    "fmt"
    "gitee.com/MM-Q/verman"
)

func main() {
    // 直接调用全局函数获取版本信息
    fmt.Println("版本信息:", verman.Version())
    fmt.Println("简洁格式:", verman.Simple())
    fmt.Println("完整信息:", verman.Complete())
}
```

### 高级用法 - 编译时注入

```bash
# 使用 ldflags 注入版本信息
go build -ldflags "
-X 'gitee.com/MM-Q/verman.AppName=myapp' 
-X 'gitee.com/MM-Q/verman.GitVersion=v1.2.3' 
-X 'gitee.com/MM-Q/verman.GitCommit=abc1234' 
-X 'gitee.com/MM-Q/verman.GitTreeState=clean' 
-X 'gitee.com/MM-Q/verman.GitCommitTime=2024-01-01T12:00:00Z' 
-X 'gitee.com/MM-Q/verman.BuildTime=2024-01-01T12:30:00Z'
" main.go
```

### CLI 应用示例

```go
package main

import (
    "flag"
    "fmt"
    "os"
    "gitee.com/MM-Q/verman"
)

func main() {
    var showVersion = flag.Bool("version", false, "显示版本信息")
    var format = flag.String("format", "simple", "版本信息格式 (simple|full|detail|complete)")
    flag.Parse()

    if *showVersion {
        switch *format {
        case "simple":
            fmt.Println(verman.Simple())
        case "full":
            fmt.Println(verman.Full())
        case "detail":
            fmt.Println(verman.Detail())
        case "complete":
            fmt.Println(verman.Complete())
        default:
            fmt.Println(verman.Version())
        }
        os.Exit(0)
    }

    // 应用程序主逻辑...
}
```

## 📚 API 文档概述

### 全局变量（只读）

| 变量名 | 类型 | 描述 |
|--------|------|------|
| `AppName` | `string` | 应用程序名称 |
| `GitVersion` | `string` | Git 语义化版本号 |
| `GitCommit` | `string` | Git 提交哈希值 |
| `GitTreeState` | `string` | Git 仓库状态 |
| `GitCommitTime` | `string` | Git 提交时间 |
| `BuildTime` | `string` | 构建时间 |
| `GoVersion` | `string` | Go 运行时版本 |
| `Platform` | `string` | 平台信息 |

### 全局函数

| 函数名 | 返回值 | 描述 |
|--------|--------|------|
| `Version()` | `string` | 标准版本格式 |
| `Simple()` | `string` | 简洁版本格式 |
| `Full()` | `string` | 完整版本格式 |
| `Detail()` | `string` | 详细版本格式 |
| `Complete()` | `string` | 完整信息格式 |

## 🎨 支持的输出格式

| 格式 | 示例输出 | 使用场景 |
|------|----------|----------|
| `Version()` | `myapp version v1.0.0 linux/amd64` | 标准版本显示 |
| `Simple()` | `myapp v1.0.0` | 简洁版本显示 |
| `Full()` | `myapp version v1.0.0 linux/amd64 (commit: abc1234)` | 包含提交信息 |
| `Detail()` | `myapp v1.0.0 linux/amd64 built at 2024-01-01` | 包含构建时间 |
| `Complete()` | `myapp v1.0.0 linux/amd64 (commit: abc1234, tree: clean, built: 2024-01-01, go: go1.21)` | 完整详细信息 |

## ⚙️ 配置选项

### 编译时注入变量

通过 `-ldflags -X` 可以注入以下私有变量：

```bash
-X 'gitee.com/MM-Q/verman.AppName=应用名称'
-X 'gitee.com/MM-Q/verman.GitVersion=版本号'
-X 'gitee.com/MM-Q/verman.GitCommit=提交哈希'
-X 'gitee.com/MM-Q/verman.GitTreeState=仓库状态'
-X 'gitee.com/MM-Q/verman.GitCommitTime=提交时间'
-X 'gitee.com/MM-Q/verman.BuildTime=构建时间'
```

### 默认值

如果未注入相应值，将使用以下默认值：
- 所有字符串变量：`"unknown"`
- 运行时变量：自动获取当前环境信息

## 📁 项目结构

```
verman/
├── 📄 verman.go          # 主要库文件
├── 🧪 verman_test.go     # 测试文件
├── 📖 README.md          # 项目文档
├── 📋 APIDOC.md          # API 文档
├── 📜 LICENSE            # 许可证文件
├── 📦 go.mod             # Go 模块文件
├── 📁 script/            # 构建脚本目录
│   ├── build.bat         # Windows 构建脚本
│   ├── build.sh          # Linux/macOS 构建脚本
│   └── build.py          # 跨平台 Python 构建脚本
├── 📁 example/           # 使用示例
│   └── main.go           # 示例程序
└── 📁 test/              # 测试项目
    ├── go.mod
    └── main.go
```

## 🧪 测试说明

### 运行测试

```bash
# 运行所有测试
go test -v

# 运行基准测试
go test -bench=. -benchmem

# 查看测试覆盖率
go test -cover
```

### 测试功能

- ✅ 所有版本格式输出测试
- ✅ 默认值处理测试
- ✅ 运行时信息获取测试
- ✅ 并发安全测试
- ✅ 基准性能测试

### 示例测试输出

```
=== RUN   TestVersion
--- PASS: TestVersion (0.00s)
=== RUN   TestSimple
--- PASS: TestSimple (0.00s)
=== RUN   TestFull
--- PASS: TestFull (0.00s)
=== RUN   TestDetail
--- PASS: TestDetail (0.00s)
=== RUN   TestComplete
--- PASS: TestComplete (0.00s)
=== RUN   TestDefaultValues
--- PASS: TestDefaultValues (0.00s)
=== RUN   TestRuntimeInfo
--- PASS: TestRuntimeInfo (0.00s)
PASS
```

## 📄 许可证

本项目采用 [MIT 许可证](LICENSE)。

## 🤝 贡献指南

我们欢迎所有形式的贡献！

### 如何贡献

1. 🍴 Fork 本仓库
2. 🌿 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 💾 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 📤 推送到分支 (`git push origin feature/AmazingFeature`)
5. 🔄 创建 Pull Request

### 贡献类型

- 🐛 Bug 修复
- ✨ 新功能开发
- 📚 文档改进
- 🧪 测试用例添加
- 🎨 代码优化

## 📞 联系方式和相关链接

### 🔗 仓库地址
- **主仓库**: [https://gitee.com/MM-Q/verman](https://gitee.com/MM-Q/verman)

### 📋 相关资源
- 📖 [API 文档](APIDOC.md)
- 🐛 [问题反馈](https://gitee.com/MM-Q/verman/issues)
- 💡 [功能建议](https://gitee.com/MM-Q/verman/issues)
- 📦 [发布版本](https://gitee.com/MM-Q/verman/releases)

### 👨‍💻 维护者
- **MM-Q** - *项目创建者和主要维护者*

---

<div align="center">

**⭐ 如果这个项目对你有帮助，请给它一个 Star！**

[🏠 返回仓库首页](https://gitee.com/MM-Q/verman) | [📖 查看文档](APIDOC.md) | [🐛 报告问题](https://gitee.com/MM-Q/verman/issues)

</div>