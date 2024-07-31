package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num byte = 43
	fmt.Printf("类型是： %T", num)
	fmt.Println()
	fmt.Println(unsafe.Sizeof(num))

	var num1 float32 = 3.14
	fmt.Println(num1)
	var num2 float32 = -3.14
	fmt.Println(num2)
	var num3 float32 = 314e-2
	fmt.Println(num3)
	var num4 float32 = 314e-2
	fmt.Println(num4)
	var num5 float32 = 314e3
	fmt.Println(num5)
	var num6 float32 = 314e+3
	fmt.Println(num6)

	var ding = 'A'
	fmt.Printf("类型是： %T", ding)
	fmt.Println()
	fmt.Printf("文字是： %c", ding)
	fmt.Println()
	fmt.Println("对应的unicode是：", ding)
}
