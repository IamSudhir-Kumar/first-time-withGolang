package main

import (
	"fmt"
)

const (
	a int     = 44
	b float32 = 55.22
	c string  = "Fuck Time"
	d bool    = true
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Printf("d: %v\n", d)
	fmt.Printf("d: %T\n", d)
	fmt.Print(c, "\n")

	fmt.Printf("%v\n", b)
	fmt.Printf("%#v\n", c)

	fmt.Printf("%T\n", c)

	fmt.Printf("%\n", a)
}
