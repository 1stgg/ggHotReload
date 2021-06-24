package gg

import (
	"fmt"
	"testing"
	"time"
)

func TestDebounce(t *testing.T) {

	fc := Debounce(func() {
		fmt.Println("执行Debounce")
	}, 500)
	// timer := time.NewTicker(Ms(500))
	timer := SetInterval(func() {
		fc()
	}, 600)

	time.Sleep(Ms(2 * 1000))
	ClearInterval(timer)
	t.Errorf("excepted:%v, got:%v", 1, 2)
}
