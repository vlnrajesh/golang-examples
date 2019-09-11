package main

import "fmt"

func main() {
	const MAX int = 3
	a := []int{10, 11, 4}
	var ptr *int
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
		ptr = &a[i]
		fmt.Println(ptr)
		fmt.Println(&a[i])
	}
}
