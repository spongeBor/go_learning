package constant_test

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	wednesday
)
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}
func TestConstantTry1(t *testing.T) {
	a := 7
	t.Log(a&Readable, a&Writable, a&Executable)
}
