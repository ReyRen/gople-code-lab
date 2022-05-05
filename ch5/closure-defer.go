package main

import "fmt"

func main() {
	x, y := 1, 2
	defer func(a int) {
		/*函数体内某个变量作为defer匿名函数的参数，则在定义defer时已获得值拷贝，否则引用某个变量的地址(引用拷贝)*/
		fmt.Println("defer x, y = ", a, y) // y is closure refer
	}(x) // x is value copy

	x += 100
	y += 200

	fmt.Println(x, y)
}
