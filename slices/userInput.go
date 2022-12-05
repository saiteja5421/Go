package main

import (
	"fmt"
	"os"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// for {
	// 	fmt.Print("Enter Text: ")
	// 	// reads user input until \n by default
	// 	scanner.Scan()
	// 	// Holds the string that was scanned
	// 	text := scanner.Text()
	// 	if len(text) != 0 {
	// 		fmt.Println(text)
	// 	} else {
	// 		// exit if user entered an empty string
	// 		break
	// 	}

	// }

	// // handle error
	// if scanner.Err() != nil {
	// 	fmt.Println("Error: ", scanner.Err())
	// }

	// to get slice from userinput
	// var s = []string{}
	// for {
	// 	reader := bufio.NewScanner(os.Stdin)
	// 	fmt.Println("Enter the names: ")
	// 	reader.Scan()
	// 	mystr := reader.Text()
	// 	if mystr == " " {
	// 		break
	// 	}
	// 	s = append(s, mystr)
	// 	fmt.Println(s)
	// 	_ = s
	// }

	// to get array from userinput
	// var m int
	// var n int
	// fmt.Print("Enter the size of array: ")
	// fmt.Scanln(&m,&n)
	// fmt.Println(m)
	// var arr = make([]int, m)
	// for i := 0; i < m; i++ {
	// 	fmt.Printf("Enter %dth element: ", i)
	// 	fmt.Scanf("%d", &arr[i])
	// }
	// fmt.Println(arr)

	args := os.Args[1:]
	fmt.Println(args)

	//declaring and initializing sum and product of type float64
	//sum, product := 0., 1.
	var n []string
	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 and 10 numbers!")
	} else {

		for _, v := range args {
			//num, err := strconv.ParseFloat(v, 64)
			// if err != nil {
			// 	//fmt.Println(err)
			// 	continue //if it can't convert string to float64, continue with the next number
			// }
			// sum += num
			// product *= num
			n = append(n, v)
		}
	}
	//n = append(n, args...)
	fmt.Printf("%T", n)
	fmt.Println(n)
}
