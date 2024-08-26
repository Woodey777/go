package main

import "fmt"

func main() {
	a, b := division(6)
	fmt.Print(a, b)

}

func division(n float64) (float64, float64)  {
	return n / 2, n / 3
}

