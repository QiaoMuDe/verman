package verman_test

import (
	"testing"

	"gitee.com/MM-Q/verman"
)

// TestGetVersionInfo 测试 GetVersionInfo 函数，验证版本信息是否正确返回。
// 该测试用例主要关注在默认情况下，版本信息是否符合预期的未知状态。
func TestGetVersionInfo(t *testing.T) {
	// 定义测试用例结构体，包含测试名称和期望的版本信息。
	tests := []struct {
		name string
		want verman.VerMan
	}{
		// 测试用例 "default values" 检查在没有特别设置时，版本信息的默认值。
		{
			name: "default values",
			want: verman.VerMan{
				AppName:       "unknown",
				GitVersion:    "unknown",
				GitCommit:     "unknown",
				GitTreeState:  "unknown",
				GitCommitTime: "unknown",
				BuildTime:     "unknown",
				GoVersion:     "go1.24.2",
				Platform:      "windows/amd64",
			},
		},
	}

	// 遍历测试用例数组，对每个测试用例执行测试。
	for _, tt := range tests {
		// 使用测试用例的名称来运行子测试。
		t.Run(tt.name, func(t *testing.T) {
			// 调用 Get 函数获取当前的版本信息。
			got := verman.Get()
			// 比较获取到的版本信息与期望的版本信息是否一致。
			if *got != tt.want {
				// 如果不一致，输出错误信息。
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestPrintVersion 测试不同格式下打印版本信息的功能
// 该测试函数通过模拟不同格式的输出来验证PrintVersion函数的行为
func TestPrintVersion(t *testing.T) {
	// 定义测试用例结构体，包含测试名称、格式和预期输出
	tests := []struct {
		name     string
		format   string
		expected string
	}{
		// 测试用例：json格式
		{"json format", "json", "{\"gitVersion\":\"unknown\"}"},
		// 测试用例：文本格式
		{"text format", "text", "Git Version: unknown"},
		// 测试用例：简单格式
		{"simple format", "simple", "unknown"},
	}

	// 遍历测试用例
	for _, tt := range tests {
		// 使用测试用例的名称来运行子测试
		t.Run(tt.name, func(t *testing.T) {
			// 这里需要mock输出函数来验证输出内容
			// 实际测试中可能需要更复杂的验证
			verman.Get().PrintVersion(tt.format)
		})
	}
}
