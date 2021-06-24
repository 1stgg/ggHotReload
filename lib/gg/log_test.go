package gg

import (
	"testing"
)

func TestLog(t *testing.T) {

	Log("abc")

	t.Errorf("excepted:%v, got:%v", 1, 2)
}
