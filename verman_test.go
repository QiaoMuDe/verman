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

	// 1. Short 格式
	fmt.Println("1. Short() - 极简格式:")
	fmt.Printf("%s\n", info.Short())
	fmt.Println()
	fmt.Println()

	// 2. Standard 格式
	fmt.Println("2. Standard() - 标准格式:")
	fmt.Printf("%s\n", info.Standard())
	fmt.Println()
	fmt.Println()

	// 3. Full 格式
	fmt.Println("3. Full() - 完整格式:")
	fmt.Printf("%s\n", info.Full())
	fmt.Println()
	fmt.Println()

	// 4. Lines 格式
	fmt.Println("4. Lines() - 分行格式:")
	fmt.Printf("%s\n", info.Lines())
	fmt.Println()
	fmt.Println()

	// 5. Table 格式
	fmt.Println("5. Table() - 表格格式:")
	fmt.Printf("%s\n", info.Table())
	fmt.Println()
	fmt.Println()

	// 6. JSON 格式
	fmt.Println("6. JSON() - JSON格式:")
	fmt.Printf("%s\n", info.JSON())
	fmt.Println()
	fmt.Println()

	// 7. KV 格式
	fmt.Println("7. KV() - 键值对格式:")
	fmt.Printf("%s\n", info.KV())
	fmt.Println()
	fmt.Println()

	// 8. Banner 格式
	fmt.Println("8. Banner() - 横幅格式:")
	fmt.Printf("%s\n", info.Banner())
	fmt.Println()
	fmt.Println()

	// 9. CSV 格式
	fmt.Println("9. CSV() - 逗号分隔格式:")
	fmt.Printf("%s\n", info.CSV())
	fmt.Println()
	fmt.Println()

	// 10. URI 格式
	fmt.Println("10. URI() - URI格式:")
	fmt.Printf("%s\n", info.URI())
	fmt.Println()
	fmt.Println()

	fmt.Println("=== 测试完成 ===")

	// 验证所有方法都返回非空字符串
	formats := map[string]string{
		"Short":    info.Short(),
		"Standard": info.Standard(),
		"Full":     info.Full(),
		"Lines":    info.Lines(),
		"Table":    info.Table(),
		"JSON":     info.JSON(),
		"KV":       info.KV(),
		"Banner":   info.Banner(),
		"CSV":      info.CSV(),
		"URI":      info.URI(),
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
	fmt.Printf("V.Short(): %s\n", V.Short())
	fmt.Printf("V.Standard(): %s\n", V.Standard())
	fmt.Println()
}

// TestEdgeCases 测试边界情况
func TestEdgeCases(t *testing.T) {
	// 测试空值情况
	emptyInfo := &Info{}

	fmt.Println("=== 边界情况测试 ===")
	fmt.Println("空 Info 结构体的输出:")
	fmt.Printf("Short(): '%s'\n", emptyInfo.Short())
	fmt.Printf("Standard(): '%s'\n", emptyInfo.Standard())
	fmt.Println()

	// 测试部分字段为空的情况
	partialInfo := &Info{
		AppName:    "TestApp",
		GitVersion: "v1.0.0",
		Platform:   "windows/amd64",
	}

	fmt.Println("部分字段填充的 Info 结构体:")
	fmt.Printf("Short(): %s\n", partialInfo.Short())
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

	b.Run("Short", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			info.Short()
		}
	})

	b.Run("Standard", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			info.Standard()
		}
	})

	b.Run("Full", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			info.Full()
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
