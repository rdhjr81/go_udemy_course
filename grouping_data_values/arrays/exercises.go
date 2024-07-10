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
