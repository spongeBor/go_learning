package type_test

import "testing"

type MyInt int64

// 不支持任何的隐式转换
func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64 = int64(a)
	t.Log(a, b)
	var c MyInt = MyInt(a)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
