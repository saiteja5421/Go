package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func cube(a float64) float64 {
// 	return a * a * a
// }

// func f1(n uint) (int, int) {
// 	f := 1
// 	s := 0
// 	for i := 1; i <= int(n); i++ {
// 		f *= i
// 		s += i
// 	}
// 	return f, s
// }
// func searchItem(s []string, b string) bool {
// 	g := false
// 	for _, v := range s {
// 		if strings.EqualFold(v, b) {
// 			g = true
// 		}
// 	}
// 	return g
// }

// func print(msg string) {
// 	fmt.Println(msg)
// }

func main() {
	// var sai float64
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	input, _ := reader.ReadString('\n')
	// i, _ := strconv.Atoi(input)
	trammed := strings.TrimSpace(input)
	fmt.Println(trammed)
	// sai1, _ := fmt.Scan(&sai)
	// fmt.Println(sai1 * 3)
	// b := cube(3)
	// fmt.Println(b)
	// c, d := f1(4)
	// fmt.Println(c, d)
	// m := searchItem([]string{"lion", "cat", "bear"}, "Bear")
	// fmt.Println(m)
	// defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Printf("Welcome %v", trammed)
}
