package padding

import "fmt"

// 左にスペース
func Padding1(str string) string {
	return fmt.Sprintf("%8s", str)
}

// 右にスペース
func Padding2(str string) string {
	return fmt.Sprintf("%-8s", str)
}
