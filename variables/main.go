package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string

	x := 20
	y := 15.5
	z := "Gopher!"

	fmt.Println(a, b, c, d, x, y, z)
}

func exercise2() {
	// multiple variable declaration
	var (
		a int
		b float64
		c bool
		d string
	)
	// multiple short variable decalaration
	x, y, z := 20, 15.5, "Gopher!"
	//fmt.Println(a, b, c, d, x, y, z)
	_, _, _, _, _, _, _ = a, b, c, d, y, x, z
}
func ex3() {
	var a float64 = 7.1

	x, y := true, 3.7

	a, z := 5.5, false

	_, _, _, _ = a, x, y, z
	// string should be in double qoutes
	//name := 'Golang' --> Error
	name := "Golang"
	fmt.Println(name)
}
