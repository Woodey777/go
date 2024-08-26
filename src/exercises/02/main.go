package main

import "fmt"

const (
	Male string = "male"
	Female 	    = "female"
) 

type Person struct {
	name string
	surname string
	age int
	height float64
	hands int
	sex string
}

func main() {
	var pers = Person{"nik", "amel", 22, 1.73, 2, Male}
	fmt.Print(pers)
}
