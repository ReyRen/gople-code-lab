package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) {
	fmt.Printf("%p\n", &base)
	add := func(i int) int {
		fmt.Printf("%p\n", &base)
		base += i
		return base
	}

	sub := func(i int) int {
		fmt.Printf("%p\n", &base)
		base -= i
		return base
	}

	return add, sub
}

func main() {
	f1, f2 := calc(100)

	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
}
