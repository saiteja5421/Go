package main

import "fmt"
func findvowels(s string) []string {
	str := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' || s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'O' || s[i] == 'U' {
			str = append(str, string(s[i]))
		}
     
	}
	return str
}
        
func main() {
	x := [5]int{1, 2, 3, 4, 5}
	for _, v := range x {
		fmt.Printf("%T %d\n", v, v)
	}
    a := findvowels("hello")
	fmt.Println(a)
}
