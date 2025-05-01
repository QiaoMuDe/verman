# verman - Go版本信息管理库

[![Go Report Card](https://goreportcard.com/badge/github.com/your-username/verman)](https://goreportcard.com/report/github.com/your-username/verman)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

verman是一个轻量级的Go语言版本信息管理库，用于在编译时注入和获取应用程序的版本信息。

## 功能特性

- 编译时自动注入Git版本信息
- 支持多种格式输出版本信息(JSON/Text/Simple)
- 自动获取Go运行时版本和平台信息
- 简单易用的API接口

## 安装

```bash
go get gitee.com/MM-Q/verman
```

## 使用方法

### 构建脚本

项目提供了三个构建脚本，分别适用于不同平台环境：

1. **build.bat** - Windows批处理脚本
   - 自动获取Git版本信息
   - 自动注入到Go编译参数中
   - 输出构建成功信息和版本详情

2. **build.sh** - Linux Shell脚本
   - 自动获取Git仓库状态
   - 自动格式化时间信息
   - 提供详细的构建日志

3. **build.py** - 跨平台Python脚本
   - 兼容Windows/Linux/macOS
   - 提供统一的构建体验
   - 支持自定义输出文件名

使用示例：
```bash
# Windows
build.bat

# Linux/macOS
./build.sh

# 跨平台
python build.py
```

### 编译时注入版本信息

```bash
go build -ldflags "-X 'gitee.com/MM-Q/verman.gitVersion=v1.0.0' \
  -X 'gitee.com/MM-Q/verman.gitCommit=abc1234' \
  -X 'gitee.com/MM-Q/verman.gitTreeState=clean' \
  -X 'gitee.com/MM-Q/verman.gitCommitTime=2024-01-01T12:00:00Z' \
  -X 'gitee.com/MM-Q/verman.buildTime=2024-01-01T12:00:00Z'" main.go
```

### 代码示例

```go
package main

import (
	"fmt"
	"gitee.com/MM-Q/verman"
)

func main() {
	// 获取版本信息结构体
	versionInfo := verman.Get()
	
	// 以JSON格式打印版本信息
	versionInfo.PrintVersion("json")
	
	// 以文本格式打印版本信息
	versionInfo.PrintVersion("text")
	
	// 以简洁格式打印版本信息
	versionInfo.PrintVersion("simple")
}
```

## API文档

### `Get() VerMan`

返回包含所有版本信息的VerMan结构体。

### `PrintVersion(format string)`

根据指定格式打印版本信息，支持格式："json", "text", "simple"。

### `SprintVersion(format string) (string, error)`

根据指定格式返回版本信息字符串，支持格式："json", "text", "simple"。
