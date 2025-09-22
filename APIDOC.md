# package verman

**import** `"gitee.com/MM-Q/verman"`

## VARIABLES

```go
var (
	AppName       string // 应用程序名称
	GitVersion    string // Git 语义化版本号(如 v1.0.0)
	GitCommit     string // Git 提交哈希值(如 abc1234)
	GitTreeState  string // Git 仓库状态(如 clean, dirty)
	GitCommitTime string // Git 提交时间(如 2024-01-01T12:00:00Z)
	BuildTime     string // 构建时间(如 2024-01-01T12:00:00Z)
)
```

全局版本信息变量，在编译时注入

```go
var (
	GoVersion string // Go 运行时版本(如 go1.19)
	Platform  string // 平台信息(如 linux/amd64)
)
```

运行时信息变量，自动获取

## FUNCTIONS

### func Complete() string

Complete 返回包含所有信息的完整字符串

### func Detail() string

Detail 返回格式为"程序名 v1.0.0 linux/amd64 built at 2024-01-01"的字符串

### func Full() string

Full 返回格式为"程序名 version 版本号 平台/架构 (commit: abc1234)"的字符串

### func Simple() string

Simple 返回格式为"程序名 v1.0.0"的字符串

### func Version() string

Version 返回格式为"程序名 version 版本号 平台/架构"的字符串

