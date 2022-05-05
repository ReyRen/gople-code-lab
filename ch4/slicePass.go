package main

import "fmt"

func main() {
	//slices := make([]int, 0)
	slices := []int{1, 3}
	fmt.Printf("main slices=%v main slicesP=%p main &slices=%p\n", slices, slices, &slices)
	changeSlice(slices)
	fmt.Printf("changeSlice done slices=%v changeSlice done slicesP=%p changeSlice done &slices=%p\n", slices, slices, &slices)
	appendSlice(slices)
	fmt.Printf("appendSlice done slices=%v appendSlice done slicesP=%p appendSlice done &slices=%p\n", slices, slices, &slices)
	/*change(s)
	fmt.Printf("change done s=%v change done &s=%p\n", s, &s)
	change2(s)
	fmt.Printf("change2 done s=%v change2 done &s=%p\n", s, &s)*/
}

func changeSlice(slices []int) {
	fmt.Printf("changeSlice before slices=%v changeSlice before slicesP=%p changeSlice before &slices=%p\n", slices, slices, &slices)
	slices[0] = 2
	fmt.Printf("changeSlice after slices=%v changeSlice after slicesP=%p changeSlice after &slices=%p\n", slices, slices, &slices)
}

func appendSlice(slices []int) {
	fmt.Printf("appendSlice before slices=%v appendSlice before slicesP=%p appendSlice before &slices=%p\n", slices, slices, &slices)
	slices = append(slices, 4)
	fmt.Printf("appendSlice after slices=%v appendSlice after slicesP=%p appendSlice after &slices=%p\n", slices, slices, &slices)
}

/*
	因为Go函数传递默认是值拷贝，将slice变量传入append函数相当于传了原slice变量的一个副本，注意不是拷贝底层数组，因为slice变量并不是数组，
	它仅仅是存储了底层数组的一些信息。
*/
