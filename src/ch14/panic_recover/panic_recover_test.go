package panic_recover

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		fmt.Println("Finally!")
	}()
	fmt.Println("Start")
	// panic(errors.New("Something wrong!")) // 执行defer
	os.Exit(-1) // 不执行defer
}

func TestPanicVxExi2(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
}
