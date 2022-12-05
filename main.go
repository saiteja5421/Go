package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	id   int
}

func main() {
	var b = []string{"a", "b"}
	fmt.Println(b)
	a := "aaBcd"
	for i := 0; i < len(a); i++ {
		if string(a[i]) == strings.ToUpper(string(a[i])) {
			fmt.Println(a[0:i-1] + a[i+1:])
		}
	}

}
