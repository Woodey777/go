package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Printf("%T %T %T\n", 5, "sdf", 3.45)
	fmt.Print(reflect.TypeOf(5), reflect.ValueOf("sdf").Kind(), reflect.ValueOf(3.45).Kind())
}
