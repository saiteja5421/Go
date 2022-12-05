package main

import (
	"fmt"
	"time"
)

// define a struct 'Student'
type Student struct {
	id      int
	name    string
	grade   int
	mark1   float32
	mark2   float32
	created time.Time
}

func (s Student) show() {
	// displays student details
	fmt.Println("----------------")
	fmt.Printf("Id: %d\n", s.id)
	fmt.Printf("Name: %s\n", s.name)
	fmt.Printf("Grade: %d\n", s.grade)
	fmt.Printf("Mark1: %.2f\n", s.mark1)
	fmt.Printf("Mark2: %.2f\n", s.mark2)
	fmt.Printf("Created: %s\n", s.created.String())
}
func main() {
	student1 := Student{
		id:    123,
		name:  "sai",
		grade: 9,
	}
	student1.show()
	student2 := Student{
		id:    12345,
		name:  "saiteja",
		grade: 6,
	}
	student2.show()
	type person struct {
		firstname        string
		lastname         string
		favoriteicecream []string
	}

	s1 := person{firstname: "saiteja", lastname: "reddy", favoriteicecream: []string{"chocolate", "vanilla", "strawberry"}}
	for _, v := range s1.favoriteicecream {
		fmt.Println(v)
	}

	m := map[string]person{"sai": s1}
	fmt.Println(m)

	type vehicle struct {
		doors, color string
	}
	truck := vehicle{doors: "FourWheel", color: "blue"}
	fmt.Println(truck)

}
