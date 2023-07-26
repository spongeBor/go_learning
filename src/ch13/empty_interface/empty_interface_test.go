package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	// 1. 空接口可以表示任何类型
	// 2. 通过断言来将空接口转换为定制类型
	if i, ok := p.(int); ok {
		fmt.Println("Integer", i)
		return
	}

	if s, ok := p.(string); ok {
		fmt.Println("string", s)
		return
	}
	fmt.Println("Unknow Type")
}

func DoSomething2(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")

	DoSomething2(5)
	DoSomething2("5")
}
