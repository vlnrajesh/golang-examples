package ArrayExamples

import (
	StringExamples "github.com/vlnrajesh/golang-examples/string_examples"
)

// ReverseIntArray reverse given array's each element and returns reversed array
func ReverseIntArray(x []int) []int {
	y := make([]int, len(x))
	for i, j := len(x)-1, 0; i >= 0; i, j = i-1, j+1 {
		y[j] = x[i]
	}
	return y
}

//ArrayToString returns given array into comma seperated string
func ArrayToString(x []int) string {
	var s string
	for i := 0; i < len(x); i++ {
		if x[i] == x[len(x)-1] {
			s += StringExamples.IntToString(x[i])
		} else {
			s += StringExamples.IntToString(x[i]) + ","
		}
	}
	return s
}

//PopIntArray to test Call by value or reference
func PopIntArray(a *[]int) int {
	var x int
	y := len(*a)
	x, *a = *a[y-1], *a[:y-1]
	return x
}
