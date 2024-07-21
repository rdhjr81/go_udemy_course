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

func Ex54() {
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

	personMap := map[string]person{}
	personMap[p1.lastName] = p1
	personMap[p2.lastName] = p2

	for _, v := range personMap {
		fmt.Println(v.firstName, " ", v.lastName)
		fmt.Println("flavors: ")
		for _, v := range v.favoriteFlavors {
			fmt.Print(v + " ")
		}
		fmt.Println()
	}
}

func Ex55() {
	type engine struct {
		electric bool
	}

	type vehicle struct {
		engine
		make  string
		model string
		doors int
		color string
	}

	prius := vehicle{
		engine: engine{electric: true},
		make:   "toyota",
		model:  "prius",
		doors:  4,
		color:  "red",
	}

	f150 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "hyundia",
		model: "ion",
		doors: 4,
		color: "silver",
	}
	fmt.Println(prius)
	fmt.Println(f150)

}

func Ex56() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "arnold",
		friends: map[string]int{
			"alex": 43,
			"joe":  23,
		},
		favDrinks: []string{"beer", "light beer"},
	}

	fmt.Println(p1)

	p2 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "pete",
		friends: map[string]int{
			"josh": 45,
			"joe":  78,
		},
		favDrinks: []string{"pina colada", "rose"},
	}

	fmt.Println(p2)
}
func main() {
	// Ex53()
	// Ex54()
	// Ex55()
	Ex56()
}
