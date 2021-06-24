package gg

import (
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
