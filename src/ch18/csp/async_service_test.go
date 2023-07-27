package async_service

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

// async
func AsyncService() chan string {
	// retCh := make(chan string) // 普通chan会阻塞执行
	retCh := make(chan string, 1) // buffer chan不会阻塞执行
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}
