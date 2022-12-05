package main

import "fmt"

func main() {
	var m1 map[string]bool
	m2 := map[int]string{1: "Abba", 3: "teja", 4: "reddy"}
	fmt.Printf("%#v\n", m1)
	m2[2] = "sai"
	//prints map
	fmt.Println(m2)
	//prints key 3---> value
	fmt.Println(m2[3])

	//out of range don't get error
	fmt.Println(m2[100])

	//iteration over map
	m4 := map[int]bool{1: true, 2: false, 3: false}

	for i, v := range m4 {
		fmt.Printf("key : %#v , value : %#v\n", i, v)
	}

	person := map[string][]string{"bond_james": {"shaken,not stirredd", "saiteja"}, "moneypenny": {"james bond"}}
	fmt.Println(person)

	delete(person, "bond_james")
	fmt.Println(person)

}
