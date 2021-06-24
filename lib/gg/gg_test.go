package gg

import (
	"testing"
)

func TestTest(t *testing.T) {

	Test()

	t.Errorf("excepted:%v, got:%v", 1, 2)
}
