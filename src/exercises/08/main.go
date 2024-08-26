package main

import (
	"fmt"
	"time"
	"math/rand"
)

func getRandomNumber(l, r float64) float64 {
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	return gen.Float64() * (r - l) + l
} 

func main() {
	fmt.Println(getRandomNumber(10, 20))
	fmt.Println(getRandomNumber(-10, 2))
	fmt.Println(getRandomNumber(-16, -2))
}
