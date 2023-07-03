package try_test

import "testing"

func TestFirstTry(t *testing.T) {
	t.Log("My first try!")
}

func TestExchange(t *testing.T) {
	a := 1
	b := 1
	tmp := a
	a = b
	b = tmp
	t.Log(a, b)
}
