package main

import "fmt"

func main() {
	s := "abcdefg"
	//tmp := make([]int, 0)
	fmt.Printf("main s=%v main &s=%p\n", s, &s)
	change(s)
	fmt.Printf("change done s=%v change done &s=%p\n", s, &s)
	change2(s)
	fmt.Printf("change2 done s=%v change2 done &s=%p\n", s, &s)
}

func change(s string) {
	fmt.Printf("change before s=%v change before &s=%p\n", s, &s)
	s = "abcdefgcccc"
	fmt.Printf("change after s=%v change after &s=%p\n", s, &s)
}

func change2(g string) {
	fmt.Printf("change2 before g=%v change2 before &g=%p\n", g, &g)
	g = "abcdefghgggg"
	fmt.Printf("change2 after g=%v change2 after &g=%p\n", g, &g)
}
