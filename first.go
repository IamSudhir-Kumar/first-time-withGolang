package main

import (
	"fmt"
)

func add(a int, b int) int {
	Sum := a + b

	return Sum
}
func subs(a int, b int) int {
	Sub := a - b

	return Sub
}
func multi(a float32, b float32) {
	Mult := a * b

	// return Mult

	fmt.Println(Mult)
}

func pri(a string) {
	fmt.Println(a)
}

func main() {
	var gio float32 = 44.55

	fmt.Println(add(4, 5))
	fmt.Println(subs(8, 10))
	multi(3.55, 4.454)
	pri("Time is everything")
	fmt.Println(gio)
}
