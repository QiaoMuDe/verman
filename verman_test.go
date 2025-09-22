package verman

import (
	"fmt"
	"testing"
)

// TestAllFormats 测试所有版本信息格式
func TestAllFormats(t *testing.T) {
	// 创建模拟的 Info 结构体并填充测试数据
	info := &Info{
		AppName:       "MyApp",
		GitVersion:    "v2.1.0",
		GitCommit:     "a1b2c3d4e5f6",
		GitTreeState:  "clean",
		GitCommitTime: "2024-03-15T14:30:00Z",
		BuildTime:     "2024-03-15T15:00:00Z",
		GoVersion:     "go1.22.1",
		Platform:      "linux/amd64",
	}

	fmt.Println("=== 版本信息格式展示 ===")
	fmt.Println()

	// 1. Simple 格式
	fmt.Println("1. Simple() - 简洁格式:")
	fmt.Printf("%s\n", info.Simple())
	fmt.Println()
	fmt.Println()

	// 2. Version 格式
	fmt.Println("2. Version() - 标准版本格式:")
	fmt.Printf("%s\n", info.Version())
	fmt.Println()
	fmt.Println()

	// 3. Full 格式
	fmt.Println("3. Full() - 完整版本格式:")
	fmt.Printf("%s\n", info.Full())
	fmt.Println()
	fmt.Println()

	// 4. Detail 格式
	fmt.Println("4. Detail() - 详细信息格式:")
	fmt.Printf("%s\n", info.Detail())
	fmt.Println()
	fmt.Println()

	// 5. Complete 格式
	fmt.Println("5. Complete() - 完整信息格式:")
	fmt.Printf("%s\n", info.Complete())
	fmt.Println()
	fmt.Println()

	// 6. Banner 格式 (多行)
	fmt.Println("6. Banner() - 横幅格式 (多行):")
	fmt.Printf("%s\n", info.Banner())
	fmt.Println()
	fmt.Println()

	// 7. Build 格式 (多行)
	fmt.Println("7. Build() - 构建信息格式 (多行):")
	fmt.Printf("%s\n", info.Build())
	fmt.Println()
	fmt.Println()

	// 8. Git 格式 (多行)
	fmt.Println("8. Git() - Git信息格式 (多行):")
	fmt.Printf("%s\n", info.Git())
	fmt.Println()
	fmt.Println()

	// 9. Table 格式 (多行)
	fmt.Println("9. Table() - 表格格式 (多行):")
	fmt.Printf("%s\n", info.Table())
	fmt.Println()
	fmt.Println()

	// 10. JSON 格式 (多行)
	fmt.Println("10. JSON() - JSON格式 (多行):")
	fmt.Printf("%s\n", info.JSON())
	fmt.Println()
	fmt.Println()

	fmt.Println("=== 测试完成 ===")

	// 验证所有方法都返回非空字符串
	formats := map[string]string{
		"Simple":   info.Simple(),
		"Version":  info.Version(),
		"Full":     info.Full(),
		"Detail":   info.Detail(),
		"Complete": info.Complete(),
		"Banner":   info.Banner(),
		"Build":    info.Build(),
		"Git":      info.Git(),
		"Table":    info.Table(),
		"JSON":     info.JSON(),
	}

	for name, result := range formats {
		if result == "" {
			t.Errorf("%s() returned empty string", name)
		}
		if len(result) < 5 {
			t.Errorf("%s() returned too short string: %q", name, result)
		}
	}
}

// TestGlobalInstance 测试全局实例 V
func TestGlobalInstance(t *testing.T) {
	if V == nil {
		t.Fatal("Global instance V should not be nil")
	}

	fmt.Println("\n=== 全局实例 V 的版本信息 ===")
	fmt.Printf("AppName: %s\n", V.AppName)
	fmt.Printf("GitVersion: %s\n", V.GitVersion)
	fmt.Printf("Platform: %s\n", V.Platform)
	fmt.Printf("GoVersion: %s\n", V.GoVersion)
	fmt.Println()

	fmt.Println("使用全局实例 V 调用方法:")
	fmt.Printf("V.Simple(): %s\n", V.Simple())
	fmt.Printf("V.Version(): %s\n", V.Version())
	fmt.Println()
}

// TestEdgeCases 测试边界情况
func TestEdgeCases(t *testing.T) {
	// 测试空值情况
	emptyInfo := &Info{}

	fmt.Println("=== 边界情况测试 ===")
	fmt.Println("空 Info 结构体的输出:")
	fmt.Printf("Simple(): '%s'\n", emptyInfo.Simple())
	fmt.Printf("Version(): '%s'\n", emptyInfo.Version())
	fmt.Println()

	// 测试部分字段为空的情况
	partialInfo := &Info{
		AppName:    "TestApp",
		GitVersion: "v1.0.0",
		Platform:   "windows/amd64",
	}

	fmt.Println("部分字段填充的 Info 结构体:")
	fmt.Printf("Simple(): %s\n", partialInfo.Simple())
	fmt.Printf("Full(): %s\n", partialInfo.Full())
	fmt.Printf("Table():\n%s\n", partialInfo.Table())
}

// BenchmarkAllMethods 性能基准测试
func BenchmarkAllMethods(b *testing.B) {
	info := &Info{
		AppName:       "BenchApp",
		GitVersion:    "v1.0.0",
		GitCommit:     "abc123",
		GitTreeState:  "clean",
		GitCommitTime: "2024-01-01T12:00:00Z",
		BuildTime:     "2024-01-01T12:30:00Z",
		GoVersion:     "go1.22.0",
		Platform:      "linux/amd64",
	}

	b.Run("Simple", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			info.Simple()
		}
	})

	b.Run("Version", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			info.Version()
		}
	})

	b.Run("Complete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			info.Complete()
		}
	})

	b.Run("JSON", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			info.JSON()
		}
	})

	b.Run("Table", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			info.Table()
		}
	})
}
