package verman

// 功能说明:
// 1. 提供编译时注入版本信息的全局变量
// 2. 提供获取不同格式版本信息的全局函数
// 3. 简洁直接的API设计，无需构造函数

/*
示例编译时注入版本信息:
go build -ldflags "-X 'gitee.com/MM-Q/verman.AppName=myapp' -X 'gitee.com/MM-Q/verman.GitVersion=v1.0.0' -X 'gitee.com/MM-Q/verman.GitCommit=abc1234' -X 'gitee.com/MM-Q/verman.GitTreeState=clean' -X 'gitee.com/MM-Q/verman.GitCommitTime=2024-01-01T12:00:00Z' -X 'gitee.com/MM-Q/verman.BuildTime=2024-01-01T12:00:00Z'" main.go
*/

import (
	"fmt"
	"runtime"
)

// 全局版本信息变量，在编译时注入
var (
	AppName       string // 应用程序名称
	GitVersion    string // Git 语义化版本号(如 v1.0.0)
	GitCommit     string // Git 提交哈希值(如 abc1234)
	GitTreeState  string // Git 仓库状态(如 clean, dirty)
	GitCommitTime string // Git 提交时间(如 2024-01-01T12:00:00Z)
	BuildTime     string // 构建时间(如 2024-01-01T12:00:00Z)
)

// 运行时信息变量，自动获取
var (
	GoVersion string // Go 运行时版本(如 go1.19)
	Platform  string // 平台信息(如 linux/amd64)
)

// 初始化函数，设置默认值和运行时信息
func init() {
	// 设置默认值
	if AppName == "" {
		AppName = "unknown"
	}
	if GitVersion == "" {
		GitVersion = "unknown"
	}
	if GitCommit == "" {
		GitCommit = "unknown"
	}
	if GitTreeState == "" {
		GitTreeState = "unknown"
	}
	if GitCommitTime == "" {
		GitCommitTime = "unknown"
	}
	if BuildTime == "" {
		BuildTime = "unknown"
	}

	// 获取运行时信息
	GoVersion = runtime.Version()
	Platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
}

// Version 返回格式为"程序名 version 版本号 平台/架构"的字符串
func Version() string {
	return fmt.Sprintf("%s version %s %s", AppName, GitVersion, Platform)
}

// Simple 返回格式为"程序名 v1.0.0"的字符串
func Simple() string {
	return fmt.Sprintf("%s %s", AppName, GitVersion)
}

// Full 返回格式为"程序名 version 版本号 平台/架构 (commit: abc1234)"的字符串
func Full() string {
	return fmt.Sprintf("%s version %s %s (commit: %s)", AppName, GitVersion, Platform, GitCommit)
}

// Detail 返回格式为"程序名 v1.0.0 linux/amd64 built at 2024-01-01"的字符串
func Detail() string {
	return fmt.Sprintf("%s %s %s built at %s", AppName, GitVersion, Platform, BuildTime)
}

// Complete 返回包含所有信息的完整字符串
func Complete() string {
	return fmt.Sprintf("%s %s %s (commit: %s, tree: %s, built: %s, go: %s)",
		AppName, GitVersion, Platform, GitCommit, GitTreeState, BuildTime, GoVersion)
}
