package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 10
	var f float64 = 6.44
	var str1 string = "101"
	var str2 string = "10.123"

	fmt.Println(float64(f))
	fmt.Println(int(i))
	// string to int
	newlint, a := strconv.ParseInt(str1, 0, 64)
	fmt.Println(newlint, a)

	//string to float
	newfloat, _ := strconv.ParseFloat(str2, 64)
	fmt.Println(newfloat)

	// int to string
	newString := fmt.Sprintf("%d", i)
	fmt.Printf("%T", newString)
	fmt.Printf("%#v\n", newString)

	//string to int
	i, err := strconv.Atoi(str1)
	fmt.Println(i)
	_ = err

	// int to string
	j := strconv.Itoa(i)
	fmt.Printf("%#v", j)
	fmt.Println(j)
}
