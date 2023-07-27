package series

import "fmt"

func init() {
	fmt.Println("init1")
}

// 可以定义多个
func init() {
	fmt.Println("init1")
}

func GetFibonacciSeries(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}

// 小写的不能在包外被访问
func square(n int) int {
	return n
}
