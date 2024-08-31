package exercises

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

var s = `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

func Ex1() {

	//marshal to json

	sJson, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sJson))
}

type Person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func Ex2() {
	sJson, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(string(sJson))
	var peopleString string
	var people []Person

	err = json.Unmarshal(sJson, &peopleString)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(peopleString), &people)
	if err != nil {
		fmt.Println(err)
	}

	for _, person := range people {
		fmt.Println(person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println("\t", saying)
		}
	}
}

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func Ex3() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}

func Ex4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

type ByAge []user

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (b ByAge) Print() {
	Print(b)
}

type ByLast []user

func (b ByLast) Len() int { return len(b) }

func (b ByLast) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByLast) Less(i, j int) bool {
	return b[i].Last < b[j].Last
}

func (b ByLast) Print() {
	Print(b)
}

func Ex5() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	ba := ByAge(users)
	sort.Sort(ba)

	fmt.Println("sorted by age")
	ba.Print()
	//sort by age

	// //sort by last
	bl := ByLast(users)
	sort.Sort(bl)

	fmt.Println("sorted by last")
	bl.Print()
}

func Print(u []user) {
	for _, person := range u {
		fmt.Println(person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println("\t", saying)
		}
	}
}
