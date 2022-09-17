package main

import (
	"fmt"

	"github.com/timothyl96/goo"
)

func main() {
	// Using var
	var a goo.String
	a = "abc"
	a = a.ToUpper()
	a = a[1:]
	fmt.Println(a) // Prints BC

	// Using new function
	b := goo.NewString("aBc")
	b = b.ToUpper()
	fmt.Println(b) // Prints ABC

	c := b.ToString()
	fmt.Println(c) // Prints ABC

	d := b.ToLower()
	fmt.Println(d) // Prints abc
}
