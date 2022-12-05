package main

import "fmt"

func main() {
	s := []float64{10.2, 5.6, 7.8}
	fmt.Println(s)
	s = append(s, 6.7)
	fmt.Println(s)
	s3 := make([]int, 5)
	fmt.Printf("%d", s3)
	fmt.Println(s[len(s)-1])
}

// package main

// import (
// 	"fmt"
// 	"os"
// )

//func main() {

	//taking the arguments in a new slice
	//fmt.Println(os.Args)
	//args := os.Args[1:]
	//fmt.Println(args[1:])

	// declaring and initializing sum and product of type float64
	// sum, product := 0., 1.

	// if len(args) < 2 || len(args) > 10 {
	// 	fmt.Println("Please enter between 2 and 10 numbers!")
	// } else {

	// 	for _, v := range args {
	// 		num, err := strconv.ParseFloat(v, 64)
	// 		if err != nil {
	// 			//fmt.Println(err)
	// 			continue //if it can't convert string to float64, continue with the next number
	// 		}
	// 		sum += num
	// 		product *= num
	// 	}
	// }

	//fmt.Printf("Sum: %v, Product: %v\n", sum, product)

	// s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// for _, v := range s {
	// 	fmt.Printf("%T-%d\n", v, v)
	// }

	// s1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// fmt.Println(s1[:5])
	// fmt.Println(s1[5:])
	// fmt.Println(s1[2:7])
	// fmt.Println(s1[1:6])

	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// x = append(x, 52)
	// fmt.Println(x)

	// x = append(x, 53, 54, 55)
	// fmt.Println(x)

	// y := []int{56, 57, 58, 59, 60}

	// x = append(x, y...)
	// fmt.Println(x)

	// z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// z = append(z[:3], z[6:]...)
	// fmt.Println(z)

	// x1 := make([]string, 5)
	// x1 = append(x1, "india")
	// fmt.Println(x1)
	// sai := []string{"sai", "saiteja", "hjgw"}
	// sai1 := []string{"saihfgk", "kyfkwesaiteja", "hrgghjgw"}
	// s2 := [][]string{sai, sai1}
	// fmt.Println(s2)
//}
