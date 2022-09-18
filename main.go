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
	// Can be slice like a builtin string
	a = a[1:]
	fmt.Println(a) // Prints BC

	// 2. Using from function (1 line)
	b := goo.FromString("aBc")
	b = b.ToUpper()
	fmt.Println(b) // Prints ABC

	c := b.ToString()
	fmt.Println(c) // Prints ABC

	d := b.ToLower()
	fmt.Println(d) // Prints abc

	// ##### Integer
	i1 := goo.FromInt32(123)
	fmt.Println(i1) // Prints 123

	var i2 goo.Int32
	// Perform arimetics using different operations
	i2 = 7
	i2 += i1
	fmt.Println(i2) // Prints 130

	// ##### Float
	f1 := goo.FromFloat64(123.232)
	fmt.Println(f1) // Prints 123
	// Convert to builtin float64
	fmt.Println(f1.ToFloat64()) // Prints 123
}
