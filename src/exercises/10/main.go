package main

import "fmt"

func main() {
	mp := make(map[int]int)

	mp[1] = 1

	fmt.Println(hasKey(mp, 1))
	fmt.Println(hasKey(mp, 3))
}

func hasKey(mp map[int]int, key int) bool {
	_, ok := mp[key]
	return ok
}