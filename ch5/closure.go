package main

import (
	"fmt"
)

func add(base int) func(int) int {
	fmt.Printf("address of base: %p\n", &base)
	f := func(i int) int {
		/*内部函数是对外部变量引用*/
		fmt.Printf("address of closure base: %p\n", &base)
		base += i
		return base
	}

	return f
}

func main() {
	a := 10
	fmt.Printf("address of a: %p\n", &a)
	t1 := add(a)
	fmt.Println(t1(1), t1(2))
	t2 := add(100)
	fmt.Println(t2(1), t2(2))
}
