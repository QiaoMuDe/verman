package verman

import (
	"runtime"
	"strings"
	"testing"
)

// 重置为默认值的辅助函数
func resetToDefaults() {
	AppName = "unknown"
	GitVersion = "unknown"
	GitCommit = "unknown"
	GitTreeState = "unknown"
	GitCommitTime = "unknown"
	BuildTime = "unknown"
}

// 设置测试数据的辅助函数
func setTestData() {
	AppName = "testapp"
	GitVersion = "v1.2.3"
	GitCommit = "abc1234"
	GitTreeState = "clean"
	GitCommitTime = "2024-01-01T12:00:00Z"
	BuildTime = "2024-01-01T12:30:00Z"
}

func TestVersion(t *testing.T) {
	// 保存原始值
	originalAppName := AppName
	originalGitVersion := GitVersion
	originalPlatform := Platform
	defer func() {
		AppName = originalAppName
		GitVersion = originalGitVersion
		Platform = originalPlatform
	}()

	// 设置测试数据
	setTestData()
	Platform = "linux/amd64"

	expected := "testapp version v1.2.3 linux/amd64"
	result := Version()

	if result != expected {
		t.Errorf("Version() = %q, want %q", result, expected)
	}
}

func TestSimple(t *testing.T) {
	// 保存原始值
	originalAppName := AppName
	originalGitVersion := GitVersion
	defer func() {
		AppName = originalAppName
		GitVersion = originalGitVersion
	}()

	// 设置测试数据
	setTestData()

	expected := "testapp v1.2.3"
	result := Simple()

	if result != expected {
		t.Errorf("Simple() = %q, want %q", result, expected)
	}
}

func TestFull(t *testing.T) {
	// 保存原始值
	originalAppName := AppName
	originalGitVersion := GitVersion
	originalGitCommit := GitCommit
	originalPlatform := Platform
	defer func() {
		AppName = originalAppName
		GitVersion = originalGitVersion
		GitCommit = originalGitCommit
		Platform = originalPlatform
	}()

	// 设置测试数据
	setTestData()
	Platform = "linux/amd64"

	expected := "testapp version v1.2.3 linux/amd64 (commit: abc1234)"
	result := Full()

	if result != expected {
		t.Errorf("Full() = %q, want %q", result, expected)
	}
}

func TestDetail(t *testing.T) {
	// 保存原始值
	originalAppName := AppName
	originalGitVersion := GitVersion
	originalBuildTime := BuildTime
	originalPlatform := Platform
	defer func() {
		AppName = originalAppName
		GitVersion = originalGitVersion
		BuildTime = originalBuildTime
		Platform = originalPlatform
	}()

	// 设置测试数据
	setTestData()
	Platform = "linux/amd64"

	expected := "testapp v1.2.3 linux/amd64 built at 2024-01-01T12:30:00Z"
	result := Detail()

	if result != expected {
		t.Errorf("Detail() = %q, want %q", result, expected)
	}
}

func TestComplete(t *testing.T) {
	// 保存原始值
	originalAppName := AppName
	originalGitVersion := GitVersion
	originalGitCommit := GitCommit
	originalGitTreeState := GitTreeState
	originalBuildTime := BuildTime
	originalPlatform := Platform
	originalGoVersion := GoVersion
	defer func() {
		AppName = originalAppName
		GitVersion = originalGitVersion
		GitCommit = originalGitCommit
		GitTreeState = originalGitTreeState
		BuildTime = originalBuildTime
		Platform = originalPlatform
		GoVersion = originalGoVersion
	}()

	// 设置测试数据
	setTestData()
	Platform = "linux/amd64"
	GoVersion = "go1.21.0"

	expected := "testapp v1.2.3 linux/amd64 (commit: abc1234, tree: clean, built: 2024-01-01T12:30:00Z, go: go1.21.0)"
	result := Complete()

	if result != expected {
		t.Errorf("Complete() = %q, want %q", result, expected)
	}
}

// 测试默认值
func TestDefaultValues(t *testing.T) {
	// 保存原始值
	originalAppName := AppName
	originalGitVersion := GitVersion
	originalGitCommit := GitCommit
	originalGitTreeState := GitTreeState
	originalGitCommitTime := GitCommitTime
	originalBuildTime := BuildTime
	defer func() {
		AppName = originalAppName
		GitVersion = originalGitVersion
		GitCommit = originalGitCommit
		GitTreeState = originalGitTreeState
		GitCommitTime = originalGitCommitTime
		BuildTime = originalBuildTime
	}()

	// 重置为默认值
	resetToDefaults()

	// 测试各个函数是否正确处理默认值
	if !strings.Contains(Version(), "unknown") {
		t.Error("Version() should contain 'unknown' when using default values")
	}

	if !strings.Contains(Simple(), "unknown") {
		t.Error("Simple() should contain 'unknown' when using default values")
	}

	if !strings.Contains(Full(), "unknown") {
		t.Error("Full() should contain 'unknown' when using default values")
	}

	if !strings.Contains(Detail(), "unknown") {
		t.Error("Detail() should contain 'unknown' when using default values")
	}

	if !strings.Contains(Complete(), "unknown") {
		t.Error("Complete() should contain 'unknown' when using default values")
	}
}

// 测试运行时信息
func TestRuntimeInfo(t *testing.T) {
	// GoVersion 应该包含 "go" 前缀
	if !strings.HasPrefix(GoVersion, "go") {
		t.Errorf("GoVersion = %q, should start with 'go'", GoVersion)
	}

	// Platform 应该包含 "/" 分隔符
	if !strings.Contains(Platform, "/") {
		t.Errorf("Platform = %q, should contain '/'", Platform)
	}

	// Platform 应该匹配当前运行时环境
	expectedPlatform := runtime.GOOS + "/" + runtime.GOARCH
	if Platform != expectedPlatform {
		t.Errorf("Platform = %q, want %q", Platform, expectedPlatform)
	}

	// GoVersion 应该匹配当前运行时版本
	expectedGoVersion := runtime.Version()
	if GoVersion != expectedGoVersion {
		t.Errorf("GoVersion = %q, want %q", GoVersion, expectedGoVersion)
	}
}

// 基准测试
func BenchmarkVersion(b *testing.B) {
	setTestData()
	for i := 0; i < b.N; i++ {
		Version()
	}
}

func BenchmarkSimple(b *testing.B) {
	setTestData()
	for i := 0; i < b.N; i++ {
		Simple()
	}
}

func BenchmarkComplete(b *testing.B) {
	setTestData()
	for i := 0; i < b.N; i++ {
		Complete()
	}
}

// 示例测试
func ExampleVersion() {
	// 注意：这个示例会使用实际的运行时值，所以输出可能会变化
	// 在实际项目中，你可以通过 -ldflags 注入具体的值
	result := Version()
	// 检查输出格式是否正确
	if !strings.Contains(result, "version") {
		panic("Version() output should contain 'version'")
	}
	// Output:
}

func ExampleSimple() {
	// 设置测试数据用于示例
	AppName = "myapp"
	GitVersion = "v1.0.0"

	result := Simple()
	if result != "myapp v1.0.0" {
		panic("unexpected output")
	}
	// Output:
}
