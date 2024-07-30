package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// fmt.Println("Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!Hello, World!")
	fmt.Println("Hello, World!Hello, World!Hello,",
		"World!Hello, World!Hello, World!Hello, World!Hello, ",
		"World!Hello, World!Hello, World!")
	// 如果字符串特别长，就采用这种方式书写。直接换行会报错
}
