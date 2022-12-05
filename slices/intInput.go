package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sample(a []int) []int {
	var arr = make([]int, len(a))
	for i, v := range a {
		arr[i] = a[v]
	}
	return arr
}

func main() {
	fmt.Printf("Enter size of your array: ")
	var size int
	fmt.Scanln(&size)
	//fmt.Printf("%", size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element: ", i)
		reader := bufio.NewScanner(os.Stdin)
		reader.Scan()
		mystr := reader.Text()
		//fmt.Printf("%T", mystr)
		myint, _ := strconv.Atoi(mystr)
		arr[i] = myint

	}
	fmt.Printf("Your array is: %v type is %T\n", arr, arr)

	m := sample(arr)
	fmt.Println(m)
}
