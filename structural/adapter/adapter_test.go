package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	var target iTarget

	// 通过适配器 Adapter，A 转化成实现 iTarget 接口定义的类型
	a := &A{}
	target = &Adaptor{a: a}
	target.f1()
	target.f2()
	target.f3()
}
