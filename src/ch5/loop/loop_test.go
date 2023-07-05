package loop_te

import "testing"

func TestWhileL(t *testing.T) {
	n := 0
	// go 语言没有while

	for n < 5 {
		t.Log(n)
		n++
	}
}
