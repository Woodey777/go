package main

import (
	"fmt"
)

func goroutGen(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func goroutSq(ch_in, ch_out chan int) {
	for num := range ch_in {
		ch_out <- num * num
	}
	close(ch_out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go goroutGen(ch1)
	go goroutSq(ch1, ch2)

	for sq_num := range ch2 {
		fmt.Println(sq_num)
	}
}