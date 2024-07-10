package arrays

import (
	"fmt"
)

func Ex42() {
	var i5 [5]int
	i5[0] = 11
	i5[1] = 22
	i5[2] = 33
	i5[3] = 44
	i5[4] = 55
	fmt.Println(i5)

	i4 := [4]int{}

	for i := 0; i < len(i4); i++ {
		i4[i] = i
	}
	fmt.Println(i4)
}
func Ex43() {
	var s10 = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for index, value := range s10 {
		fmt.Printf("index: %v, value: %v, type: %T \n", index, value, value)
	}
}

func Ex44() {
	var s10 = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	s1 := s10[:5]
	fmt.Println(s1)

	s2 := s10[5:]
	fmt.Println(s2)

	s3 := s10[2:7]
	fmt.Println(s3)

	s4 := s10[1:6]
	fmt.Println(s4)
}

func Ex45() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func Ex46() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	//goal [42, 43, 44, 48, 49, 50, 51]

	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}

func Ex47() {
	var states = make([]string, 50) //initializes a slice with len 50 and all values are 0
	states = append(states, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	states = append(states[:0], states[50:]...)

	//above works but its extra steps

	fmt.Println("lenght of states:", len(states), "capacity of states:", cap(states))
	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}

	//correct way
	var statesEmpty = make([]string, 0, 50) //initializes a slice with len 50 and w/ no values
	statesEmpty = append(statesEmpty, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println("lenght of statesEmpty:", len(statesEmpty), "capacity of statesEmpty:", cap(statesEmpty))
	for i := 0; i < len(statesEmpty); i++ {
		fmt.Println(i, statesEmpty[i])
	}
}

func Ex48() {
	sliceOfSlice := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "I'm 008."}}
	for i, s := range sliceOfSlice {
		fmt.Println("slice", s)
		for j, v := range s {
			fmt.Println(i, j, v)
		}
	}
}
