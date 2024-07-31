package main

import "fmt"

func main() {
	var num int = 18
	fmt.Println(num)

	var num2 int
	fmt.Println(num2)

	var num3 = 10

	sex := "男"

	fmt.Println(num3)
	fmt.Println(sex)

	var a, b, c = 12, "dingkaile", 44.334

	fmt.Println(a, b, c)

	d, e, f := 23, 4545, "我是一个兵"

	fmt.Println(d, e, f)

	var (
		dingkaile = "脉脉不得语"
		nanqiang  = "盈盈一水间"
	)

	fmt.Println(dingkaile, nanqiang)
}
