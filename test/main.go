package main

import (
	"fmt"

	"gitee.com/MM-Q/verman"
)

func main() {
	fmt.Println("=== verman 版本信息格式演示 ===")
	fmt.Println()

	// 1. Short 格式
	fmt.Println("1. Short() - 短格式:")
	fmt.Println(verman.V.Short())
	fmt.Println()

	// 2. Long 格式
	fmt.Println("2. Long() - 长格式:")
	fmt.Println(verman.V.Long())
	fmt.Println()

	// 3. Standard 格式
	fmt.Println("3. Standard() - 标准格式:")
	fmt.Println(verman.V.Standard())
	fmt.Println()

	// 4. Version 格式
	fmt.Println("4. Version() - 版本格式:")
	fmt.Println(verman.V.Version())
	fmt.Println()

	// 5. Full 格式
	fmt.Println("5. Full() - 完整格式:")
	fmt.Println(verman.V.Full())
	fmt.Println()

	// 6. Lines 格式
	fmt.Println("6. Lines() - 分行格式:")
	fmt.Println(verman.V.Lines())
	fmt.Println()

	// 7. Table 格式
	fmt.Println("7. Table() - 表格格式:")
	fmt.Println(verman.V.Table())
	fmt.Println()

	// 8. JSON 格式
	fmt.Println("8. JSON() - JSON格式:")
	fmt.Println(verman.V.JSON())
	fmt.Println()

	// 9. KV 格式
	fmt.Println("9. KV() - 键值对格式:")
	fmt.Println(verman.V.KV())
	fmt.Println()

	// 10. Banner 格式
	fmt.Println("10. Banner() - 横幅格式:")
	fmt.Println(verman.V.Banner())
	fmt.Println()

	// 11. CSV 格式
	fmt.Println("11. CSV() - 逗号分隔格式:")
	fmt.Println(verman.V.CSV())
	fmt.Println()

	// 12. URI 格式
	fmt.Println("12. URI() - URI格式:")
	fmt.Println(verman.V.URI())
	fmt.Println()

	fmt.Println("=== 演示完成 ===")
}
