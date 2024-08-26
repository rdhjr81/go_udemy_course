package json_stuff

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func Marshalling() []byte {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Gold",
		Last:  "Finger",
		Age:   63,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
	return bs
}

func Unmarshalling() {

	bs := Marshalling()
	fmt.Println("Unmarshalling json")
	peeps := []person{}
	err := json.Unmarshal(bs, &peeps)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("%+v", peeps)
}
