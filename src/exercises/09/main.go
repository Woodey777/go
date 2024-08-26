package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	a := "mfgah134517095aldrfgvh8h"
	var b strings.Builder
	for _, c := range a {
		if unicode.IsNumber(c) {
			b.WriteRune(c)
		}
	}
	fmt.Println(b.String())
}