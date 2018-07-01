package main

import "fmt"

func test_mul_args(a int, args ...int) (result int){
	result +=a
	for _, arg := range args{ //range 将传入index，value 两个变量
		result += arg
	}
	return
}


func main() {
	fmt.Println(test_mul_args(1, 2, 3, 4, 5))
}