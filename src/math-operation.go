package main

import "fmt"

func main() {

	var (
		a = 10
		b = 10
		c = a + b
		d = a - b
		e = a * b
		f = a / b
		g = a % b
		h = 10
	)

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	h += 10
	fmt.Println(h)

	h++
	fmt.Println(h)
}