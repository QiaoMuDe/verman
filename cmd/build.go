package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

// 可配置变量
const (
	baseOutputName    = "myapp"
	defaultOutputDir  = "output"
	defaultEntryFile  = "./main.go"
	defaultGoCompiler = "go"
)

var (
	useVendor         = flag.Bool("use-vendor", false, "是否使用vendor克隆依赖")
	useVendorInBuild  = flag.Bool("use-vendor-build", false, "构建阶段是否使用vendor目录")
	injectGitInfo     = flag.Bool("inject-git", true, "是否注入git信息")
	simpleName        = flag.Bool("simple-name", false, "是否使用简单文件名格式")
	packageZip        = flag.Bool("zip", false, "是否打包为zip文件")
	concurrency       = flag.Int("concurrency", 4, "批量构建并发线程数")
	timeout           = flag.Int("timeout", 1800, "批量构建超时时间(秒)")
)

// Git信息结构体
type gitInfo struct {
	version    string
	commit     string
	commitTime string
	status     string
}

var gitCache gitInfo

func main() {
	flag.Parse()

	// 构建前检查
	if !checkGoInstalled() || !checkGoMod() || !checkEntryFile(defaultEntryFile) {
		return
	}

	// 获取Git信息
	if *injectGitInfo {
		if err := getGitInfo(); err != nil {
			fmt.Printf("获取Git信息失败: %v\n", err)
			return
		}
	}

	// 依赖管理
	if err := runGoModTidy(); err != nil {
		fmt.Printf("go mod tidy失败: %v\n", err)
		return
	}
	if *useVendor {
		if err := runGoModVendor(); err != nil {
			fmt.Printf("go mod vendor失败: %v\n", err)
			return
		}
	}

	// 代码检查
	if err := runGoVet(); err != nil {
		fmt.Printf("go vet检查失败: %v\n", err)
		return
	}
	if err := runGofmt(); err != nil {
		fmt.Printf("gofmt格式化失败: %v\n", err)
		return
	}

	// 批量构建
	batchBuild()
}

func checkGoInstalled() bool {
	cmd := exec.Command(defaultGoCompiler, "version")
	if err := cmd.Run(); err != nil {
		fmt.Printf("错误: 未检测到Go编译器，请检查PATH\n")
		return false
	}
	fmt.Println("ok: Go编译器可用")
	return true
}

func checkGoMod() bool {
	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		fmt.Println("错误: 当前目录无go.mod文件")
		return false
	}
	fmt.Println("ok: go.mod文件存在")
	return true
}

func checkEntryFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("错误: 入口文件%v不存在\n", path)
		return false
	}
	fmt.Println("ok: 入口文件存在")
	return true
}

func runGoModTidy() error {
	cmd := exec.Command(defaultGoCompiler, "mod", "tidy")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("执行失败: %v", err)
	}
	fmt.Println("ok: go mod tidy执行成功")
	return nil
}

func runGoModVendor() error {
	cmd := exec.Command(defaultGoCompiler, "mod", "vendor")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("执行失败: %v", err)
	}
	fmt.Println("ok: go mod vendor执行成功")
	return nil
}

func runGoVet() error {
	cmd := exec.Command(defaultGoCompiler, "vet", "./...")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("检查失败: %v", err)
	}
	fmt.Println("ok: go vet检查通过")
	return nil
}

func runGofmt() error {
	cmd := exec.Command(defaultGoCompiler, "fmt", "./...")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("格式化失败: %v", err)
	}
	fmt.Println("ok: 代码格式化完成")
	return nil
}

func getGitInfo() error {
	// 获取版本
	versionCmd := exec.Command("git", "describe", "--tags", "--always", "--dirty")
	versionOut, err := versionCmd.Output()
	if err != nil {
		return fmt.Errorf("获取版本失败: %v", err)
	}
	gitCache.version = string(versionOut)

	// 获取提交哈希
	commitCmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	commitOut, err := commitCmd.Output()
	if err != nil {
		return fmt.Errorf("获取提交哈希失败: %v", err)
	}
	gitCache.commit = string(commitOut)

	// 获取提交时间
	commitTimeCmd := exec.Command("git", "log", "-1", "--format=%cd", "--date=iso")
	commitTimeOut, err := commitTimeCmd.Output()
	if err != nil {
		return fmt.Errorf("获取提交时间失败: %v", err)
	}
	gitCache.commitTime = string(commitTimeOut)

	// 获取仓库状态
	statusCmd := exec.Command("git", "diff", "--quiet")
	if err := statusCmd.Run(); err != nil {
		gitCache.status = "dirty"
	} else {
		gitCache.status = "clean"
	}

	fmt.Println("ok: Git信息获取完成")
	return nil
}

func batchBuild() {
	platforms := []string{"windows", "linux", "darwin"}
	architectures := []string{"amd64", "arm64", "386", "arm"}
	var wg sync.WaitGroup
	wg.Add(len(platforms) * len(architectures))

	for _, plat := range platforms {
		for _, arch := range architectures {
			go func(p, a string) {
				defer wg.Done()
				if p == "darwin" && (a == "386" || a == "arm") {
					fmt.Printf("跳过不支持的组合: %s/%s\n", p, a)
					return
				}
				buildSingle(p, a)
			}(plat, arch)
		}
	}

	// 等待所有任务完成或超时
	waitChan := make(chan struct {})
	go func() {
		wg.Wait()
		close(waitChan)
	}()

	select {
	case <-waitChan:
		fmt.Println("批量构建完成")
	case <-time.After(time.Duration(*timeout) * time.Second):
		fmt.Println("错误: 批量构建超时")
	}
}

func buildSingle(platform, arch string) {
	outputDir := filepath.Join(defaultOutputDir, platform, arch)
	_ = os.MkdirAll(outputDir, 0755)

	baseName := baseOutputName
	if !*simpleName {
		baseName += fmt.Sprintf("_%s_%s", platform, arch)
	}
	outputFile := filepath.Join(outputDir, baseName)
	if platform == "windows" {
		outputFile += ".exe"
	}

	ldFlags := "-s -w"
	if *injectGitInfo {
		ldFlags = fmt.Sprintf("-X 'gitee.com/MM-Q/verman.appName=%s' -X 'gitee.com/MM-Q/verman.gitVersion=%s' -X 'gitee.com/MM-Q/verman.gitCommit=%s' -X 'gitee.com/MM-Q/verman.gitCommitTime=%s' -X 'gitee.com/MM-Q/verman.buildTime=%s' -X 'gitee.com/MM-Q/verman.gitTreeState=%s' -s -w",
			baseOutputName, gitCache.version, gitCache.commit, gitCache.commitTime, time.Now().UTC().Format("2006-01-02T15:04:05Z"), gitCache.status)
	}

	cmd := exec.Command(defaultGoCompiler, "build", "-o", outputFile, "-ldflags", ldFlags, defaultEntryFile)
	if *useVendorInBuild {
		cmd.Env = append(os.Environ(), "GOFLAGS=-mod=vendor")
	}
	cmd.Env = append(os.Environ(), "GOOS="+platform, "GOARCH="+arch)

	if err := cmd.Run(); err != nil {
		fmt.Printf("错误: 构建%s/%s失败: %v\n", platform, arch, err)
		return
	}
	fmt.Printf("ok: 构建成功 %s\n", outputFile)

	if *packageZip {
		zipFile := outputFile + ".zip"
		if err := zipFile(outputFile, zipFile); err != nil {
			fmt.Printf("错误: 打包失败 %s: %v\n", zipFile, err)
			return
	}
		_ = os.Remove(outputFile)
		fmt.Printf("ok: 打包完成 %s\n", zipFile)
	}
}

func zipFile(src, dest string) error {
	zipf, err := zip.Create(dest)
	if err != nil {
		return err
	}
	defer zipf.Close()

	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Method = zip.Deflate

	writer, err := zipf.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, f)
	return err
}
"}}}