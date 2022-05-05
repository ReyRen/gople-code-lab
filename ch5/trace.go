package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	/*bigSlowOperation被调时，trace会返回一个函数值，该函数值会在bigSlowOperation退出时被调用*/
	defer trace("bigSlowOperation")() // dont forget the extra parentheses
	// ...lots of work
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s(%s)", msg, time.Since(start))
	}
}

func double(x int) (result int) {
	/*defer语句中的函数会在return语句更新返回值变量后再执行，
	在函数中定义的匿名函数可以访问该函数包括返回值变量在内的所有变量*/
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

func main() {
	//bigSlowOperation()
	//_ = double(4)
	fmt.Println(triple(4)) // 12
}
