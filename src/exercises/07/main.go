package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{37, 5, 1, 2} 
	b := []int{6, 2, 4, 37}
	fmt.Println(getCross(a, b))
	
	a = []int{1, 1, 1} 
	b = []int{1, 1, 1, 1}
	fmt.Println(getCross(a, b))
}

func getCross(a []int, b []int) []int {
	var res []int

	sort.Ints(a)
	sort.Ints(b)

	var (
		n int = 0
		m int = 0
	)

	for n != len(a) && m != len(b) {
		if a[n] == b[m] {
			res = append(res, a[n])
			n++
			m++
		} else {
			if a[n] > b[m] {
				m++
			} else {
				n++
			}
		}
	} 
	return res
}

