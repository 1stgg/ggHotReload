package gg

import (
	"fmt"
	"time"
)

// 毫秒
// var Ms = time.Millisecond
func Ms(msNum int) time.Duration {
	return time.Duration(msNum) * time.Millisecond
}

func SetInterval(callback func(), ms int) *time.Ticker {
	timer := time.NewTicker(Ms(ms))
	go func() {
		for _ = range timer.C {
			callback()
		}
	}()
	return timer
}

func ClearInterval(timer *time.Ticker) {
	timer.Stop()
}

func Debounce(fn func(), ms int) func() {
	prev := time.Unix(0, 0)

	return func() {
		curr := time.Now()
		delta := curr.Sub(prev)

		prev = curr // 可执行了之后，在刷新计时
		// fmt.Println(14, delta < Ms(ms))
		if delta < Ms(ms) {
			return
		}

		fn()
	}
}
func PrintColor(str string, mode string) {
	switch mode {
	case "warn":
		// fmt.Printf("%c[33;0m%s%c[0m\n", 0x1B, str, 0x1B)
		fmt.Printf("%c[33m%s%c[0m\n", 0x1B, str, 0x1B)
	case "success":
		fmt.Printf("%c[32m%s%c[0m\n", 0x1B, str, 0x1B)
	case "error":
		fmt.Printf("%c[31m%s%c[0m\n", 0x1B, str, 0x1B)
	}

	// 前景 背景 颜色
	// ---------------------------------------
	// 30 40 黑色
	// 31 41 红色
	// 32 42 绿色
	// 33 43 黄色
	// 34 44 蓝色
	// 35 45 紫红色
	// 36 46 青蓝色
	// 37 47 白色

	// 0 终端默认设置
	// 1 高亮显示
	// 4 使用下划线
	// 5 闪烁
	// 7 反白显示
	// 8 不可见
	// %c[{前景色（文字颜色）};{背景色}{;前景色2，可选}m
	// fmt.Printf("%c[32;32;1m%s%c[0m\n\n", 0x1B, "testPrintColor", 0x1B)
}
func Test() {
	fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, "testPrintColor", 0x1B)
}
