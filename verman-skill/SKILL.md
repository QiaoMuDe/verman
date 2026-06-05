---
name: verman-skill
description: "当用户需要为 Go 应用添加版本信息、生成版本显示代码、创建编译时注入（-ldflags -X）构建脚本、了解 verman 库的使用方式，或询问版本信息输出格式相关问题时触发此技能。触发关键词：版本信息、版本显示、构建时注入、ldflags、Go 应用版本、--version 标志、version 命令、构建脚本、verman、12 种输出格式（Short/Long/JSON/Banner/Table/CSV/URI/KV/Lines/Standard/Full/Version）以及 Go 编译时变量注入。"
---

# verman Skill — Go 版本信息管理库使用指南与代码生成

## 概述

verman 是一个零依赖的 Go 版本信息管理库，通过编译时注入（`-ldflags -X`）将 Git 版本信息嵌入二进制文件，并提供 **12 种输出格式** 的版本信息展示方法。

**核心价值**：纯标准库、全局单例 `V` 直接调用、覆盖 CLI/日志/API/启动横幅等所有版本展示场景。

**技能资源**：`scripts/` 目录预置了三个构建脚本模板（build.py / build.sh / build.bat），AI 可直接读取并按需修改后输出给用户。

---

## 一、集成步骤

### 1. 安装依赖

```bash
go get gitee.com/MM-Q/verman
go mod tidy
```

### 2. 在代码中使用

verman 通过全局实例 `V` 暴露所有方法，无需手动创建实例。在 `init()` 中会自动采集 Go 版本和平台信息。

**最简单的使用方式**：

```go
package main

import (
    "fmt"
    "gitee.com/MM-Q/verman"
)

func main() {
    // 直接通过全局实例 V 调用
    fmt.Println(verman.V.Version())
    fmt.Println(verman.V.Short())
    fmt.Println(verman.V.Full())
}
```

### 3. 编译时注入版本信息

verman 的核心机制是通过 `-ldflags -X` 在编译时注入 6 个私有变量：

```bash
go build -ldflags "
-X 'gitee.com/MM-Q/verman.appName=myapp'
-X 'gitee.com/MM-Q/verman.gitVersion=v1.2.3'
-X 'gitee.com/MM-Q/verman.gitCommit=abc1234'
-X 'gitee.com/MM-Q/verman.gitTreeState=clean'
-X 'gitee.com/MM-Q/verman.gitCommitTime=2024-01-01T12:00:00Z'
-X 'gitee.com/MM-Q/verman.buildTime=2024-01-01T12:30:00Z'
" main.go
```

> **注意**：注入的变量名是**小写**（`appName`、`gitVersion` 等），因为它们是在 verman 包中定义的私有（非导出）变量。

---

## 二、12 种输出格式速查

使用时根据具体场景选择最合适的格式：

| 方法 | 示例输出 | 推荐场景 |
|------|---------|---------|
| `Version()` | `myapp version v1.0.0 linux/amd64` | `--version` 默认输出 |
| `Short()` | `myapp v1.0.0` | 极简显示 |
| `Long()` | `myapp version v1.0.0 linux/amd64 with go1.22.1` | 详细单行日志 |
| `Standard()` | `myapp v1.0.0 (abc1234) [linux/amd64]` | 标准带 commit |
| `Full()` | `myapp v1.0.0 (abc1234) linux/amd64 built ... with go1.22.1` | 最完整信息 |
| `Lines()` | 分行文本（4 行） | 调试日志输出 |
| `Table()` | Unicode 边框表格 | 终端详表展示 |
| `JSON()` | `{"name":"...","version":"...",...}` | API 响应/配置文件 |
| `KV()` | `app=myapp version=v1.0.0 commit=...` | 结构化日志字段 |
| `Banner()` | 带 `╔╗╚╝` 边框的多行横幅 | 程序启动画面 |
| `CSV()` | `myapp,v1.0.0,abc1234,linux/amd64,...` | 数据导出 |
| `URI()` | `verman://myapp/v1.0.0?commit=...` | 统一资源标识 |

---

## 三、典型使用模式

### 模式 A：CLI 应用的 `--version` 支持

将 verman 集成到 CLI 工具的版本显示功能中，这是最常用的模式。

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
    var format = flag.String("format", "simple", "版本信息格式 (short/long/standard/full/lines/table/json/kv/banner/csv/uri)")
    flag.Parse()

    if *showVersion {
        switch *format {
        case "short":
            fmt.Println(verman.V.Short())
        case "long":
            fmt.Println(verman.V.Long())
        case "standard":
            fmt.Println(verman.V.Standard())
        case "full":
            fmt.Println(verman.V.Full())
        case "lines":
            fmt.Println(verman.V.Lines())
        case "table":
            fmt.Println(verman.V.Table())
        case "json":
            fmt.Println(verman.V.JSON())
        case "kv":
            fmt.Println(verman.V.KV())
        case "banner":
            fmt.Println(verman.V.Banner())
        case "csv":
            fmt.Println(verman.V.CSV())
        case "uri":
            fmt.Println(verman.V.URI())
        default:
            fmt.Println(verman.V.Version())
        }
        os.Exit(0)
    }

    // 应用主逻辑...
}
```

### 模式 B：程序启动横幅

适用于服务端应用，启动时打印版本横幅。

```go
package main

import (
    "fmt"
    "gitee.com/MM-Q/verman"
)

func main() {
    // 启动横幅
    fmt.Println(verman.V.Banner())
    // 或者使用表格格式
    // fmt.Println(verman.V.Table())

    // 启动服务...
}
```

### 模式 C：API 版本端点返回 JSON

适用于微服务架构，通过 API 暴露版本信息。

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "gitee.com/MM-Q/verman"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(verman.V.JSON()))
}

func main() {
    http.HandleFunc("/version", versionHandler)
    fmt.Println(verman.V.Banner())
    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}
```

### 模式 D：结构化日志输出（logrus/zerolog 等）

将版本信息以 KV 格式注入日志上下文，方便日志聚合系统提取。

```go
package main

import (
    "fmt"
    "gitee.com/MM-Q/verman"
)

func main() {
    // KV 格式适用于结构化日志
    // 输出: app=myapp version=v1.0.0 commit=abc1234 platform=linux/amd64 build=2024-01-01 go=go1.22.1
    log.WithField("version_info", verman.V.KV()).Info("application started")
}
```

---

## 四、构建脚本生成指南

> 本 Skill 的 `scripts/` 目录下已预置三个构建脚本模板（`build.py`、`build.sh`、`build.bat`），AI 可直接读取并按用户需求修改后输出。

当用户需要构建脚本来自动化注入版本信息时，按以下规则选择合适的脚本模板。

### 选择规则

| 用户平台 | 推荐脚本 | 复杂场景 |
|---------|---------|---------|
| Windows | `build.bat` 或 `build.py` | 需要交叉编译时用 `build.py` |
| Linux/macOS | `build.sh` 或 `build.py` | 需要交叉编译时用 `build.py` |
| 跨平台/CI | `build.py` | 批量交叉编译、ZIP 打包 |

### 核心构建参数

所有构建脚本的核心逻辑一致：

```
1. 通过 git describe --tags 获取版本号
2. 通过 git rev-parse --short HEAD 获取 commit 哈希
3. 通过 git status --porcelain 判断仓库状态 (clean/dirty)
4. 获取 UTC 构建时间
5. 组装 -ldflags 参数并执行 go build
```

### 推荐：使用 build.py（功能最完整）

为用户生成 build.py 时，需要修改的关键配置变量：

```python
BASE_OUTPUT_NAME = "myapp"           # ← 改成实际应用名称
DEFAULT_ENTRY_FILE = "./main.go"     # ← 改成实际入口文件路径
DEFAULT_OUTPUT_DIR = "output"        # 输出目录
```

**批量交叉编译用法**：

```bash
# 构建当前平台
python build.py -git

# 构建所有支持的平台 (windows/linux/darwin + amd64)
python build.py -batch -git

# 构建并打包为 ZIP
python build.py -batch -git -z

# 仅构建当前平台
python build.py -git -c

# 构建后自动安装到 GOPATH/bin
python build.py -git -ai
```

### 轻量方案：build.sh / build.bat

需要修改：

```bash
# build.sh
OUTPUT_FILE="myapp"          # ← 输出文件名
PROJECT_NAME="myapp"         # ← 项目名
ENTRY_FILE="main.go"         # ← 入口文件
```

```bat
:: build.bat
set OUTPUT_FILE=myapp.exe    :: ← 输出文件名
set PROJECT_NAME=myapp       :: ← 项目名
set ENTRY_FILE=main.go       :: ← 入口文件
```

---

## 五、Info 结构体说明

verman 的 `Info` 结构体存储所有版本元信息，**所有字段均为公开可读**，可以直接通过 `verman.V` 按需访问单个字段，而不只是调用格式化方法。

```go
type Info struct {
    AppName       string // 应用程序名称
    GitVersion    string // Git 语义化版本号
    GitCommit     string // Git 提交哈希值
    GitTreeState  string // Git 仓库状态 (clean/dirty)
    GitCommitTime string // Git 提交时间
    BuildTime     string // 构建时间
    GoVersion     string // Go 运行时版本 (自动采集)
    Platform      string // 平台信息 (自动采集, 如 linux/amd64)
}
```

### 灵活用法：直接访问字段

除了调用 `V.Short()`、`V.JSON()` 等格式化方法，你也可以直接读取字段自行组合，适合需要自定义输出格式的场景：

```go
// 按需读取单个字段
fmt.Printf("App: %s\n", verman.V.AppName)
fmt.Printf("Version: %s\n", verman.V.GitVersion)
fmt.Printf("Commit: %s (%s)\n", verman.V.GitCommit, verman.V.GitTreeState)
fmt.Printf("Platform: %s\n", verman.V.Platform)
fmt.Printf("Go: %s\n", verman.V.GoVersion)
fmt.Printf("Built: %s\n", verman.V.BuildTime)

// 组合使用
fmt.Printf("%s/%s @ %s", verman.V.Platform, verman.V.GoVersion, verman.V.GitCommit)
```

> 注意：`GoVersion` 和 `Platform` 由 `init()` 自动通过 `runtime` 包获取，无需通过 `-ldflags` 注入。其他 6 个字段可通过编译时注入自定义，未注入时默认值为 `"unknown"`。

---

## 六、完整集成示例

为 Go 应用添加版本信息支持的完整工作流程：

**Step 1** — 安装依赖：
```bash
go get gitee.com/MM-Q/verman
```

**Step 2** — 在 `main.go` 中添加版本显示逻辑（参考"模式 A"或"模式 B"）。

**Step 3** — 创建构建脚本（推荐使用 `build.py`，参考第四节），修改 `BASE_OUTPUT_NAME` 等配置。

**Step 4** — 构建：
```bash
python build.py -git
```

**Step 5** — 验证：
```bash
./output/myapp --version
# 或
./output/myapp --format banner
```

---

## 七、注意事项与常见问题

1. **变量名大小写**：`-ldflags -X` 注入时必须使用小写变量名（`appName` 而非 `AppName`），因为它们是包级私有变量。

2. **Go 版本兼容性**：verman 仅使用 `fmt` 和 `runtime` 标准库，理论上兼容 Go 1.18+。

3. **JSON 转义**：当前 `JSON()` 方法使用 `fmt.Sprintf` 手动拼接。如果字段值包含特殊字符（如引号），输出可能不正确。对于安全敏感场景，建议使用 `encoding/json` 替代。

4. **Table/Banner 列宽**：`Table()` 和 `Banner()` 使用固定的 16 字符列宽。如果应用名称或版本号超过此宽度，显示会错位。若需要更宽的列，建议自行格式化。

5. **运行时不可变**：`V` 在 `init()` 初始化后不会在运行时改变。如需修改字段值，可以直接赋值（字段是公开的），但不建议这样做。

6. **并发安全**：由于初始化后所有字段为只读，verman 天然并发安全，无需加锁。
