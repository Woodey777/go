package main

import "fmt"

func main() {
	fmt.Print(eq([]int{3,2,3}, []int{1,2,3}))
}

func eq(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
