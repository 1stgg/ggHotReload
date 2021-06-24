package gg

import (
	"reflect"
	"testing"
)

func TestGscanDir(t *testing.T) {
	// i := ScanDir("../")
	// fmt.Println(i)
	// if i != 10 {
	// 	t.Error("failed")
	// }
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	// 定义一个存储测试用例的切片
	tests := []test{
		{input: "./scanDir_test", sep: ":", want: []string{"./scanDir_test/", "./scanDir_test/test1/", "./scanDir_test/test1/test1-1/", "./scanDir_test/test2/"}},
	}
	// 遍历切片，逐一执行测试用例
	for _, tc := range tests {
		got := ScanDir(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("excepted:%v, got:%v", tc.want, got)
		}
	}
}
