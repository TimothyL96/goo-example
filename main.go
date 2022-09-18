package main

import (
	"fmt"

	"github.com/timothyl96/goo"
)

func main() {
	// ##### String
	// 1. Using var
	var a goo.String
	a = "abc"
	a = a.ToUpper()
	a = a[1:]
	fmt.Println(a) // Prints BC

	// 2. Using from function
	b := goo.FromString("aBc")
	b = b.ToUpper()
	fmt.Println(b) // Prints ABC

	c := b.ToString()
	fmt.Println(c) // Prints ABC

	d := b.ToLower()
	fmt.Println(d) // Prints abc

	// ##### Integer
	i1 := goo.FromInt32(123)
	fmt.Println(i1)           // Prints 123
	fmt.Println(i1.ToInt32()) // Prints 123

	// ##### Float
	f1 := goo.FromFloat64(123.232)
	fmt.Println(f1)             // Prints 123
	fmt.Println(f1.ToFloat64()) // Prints 123
}
