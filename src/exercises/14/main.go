package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("you shold provide path to file as first argument")
		return
	}

	_, e := os.Stat(os.Args[1])
	if os.IsNotExist(e) {
		fmt.Println("file ne exist")
		return
	}
	
	fmt.Println("file exist")
}