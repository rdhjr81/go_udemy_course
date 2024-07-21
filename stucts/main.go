package main

import (
	"fmt"
)

func Ex53() {
	type person struct {
		firstName       string
		lastName        string
		favoriteFlavors []string
	}

	p1 := person{
		firstName:       "arthur",
		lastName:        "smith",
		favoriteFlavors: []string{"choc", "banana"},
	}

	p2 := person{
		firstName:       "joe",
		lastName:        "momma",
		favoriteFlavors: []string{"vanilla", "orange dreamsicle"},
	}

	fmt.Println(p1.firstName, p1.lastName, "flavors:")
	for _, v := range p1.favoriteFlavors {
		fmt.Print(v)
	}
	fmt.Println()
	fmt.Println(p2.firstName, p2.lastName, "flavors:")
	for _, v := range p2.favoriteFlavors {
		fmt.Print(v)
	}
}

func Ex54(){
	type person struct {
		firstName       string
		lastName        string
		favoriteFlavors []string
	}

	p1 := person{
		firstName:       "arthur",
		lastName:        "smith",
		favoriteFlavors: []string{"choc", "banana"},
	}

	p2 := person{
		firstName:       "joe",
		lastName:        "momma",
		favoriteFlavors: []string{"vanilla", "orange dreamsicle"},
	}

	personMap := map[string, person]
}
func main() {
	Ex53()

}
