package main

import "fmt"

func main() {
	// array literal
	a1 := [2]int{1, 2}
	fmt.Println(a1)

	a2 := [...]string{"hi", "my", "name", "is", "rob"}
	fmt.Println(a2)

	var a3 [2]int
	a3[0] = 1
	a3[1] = 42

	// a1 = a3

	fmt.Println("a3[0]:", a3[0], "a3[1]:", a3[1])

	fmt.Println("a1[0]:", a1[0], "a1[1]:", a1[1])

	fmt.Println("a1 size:", len(a1))
	fmt.Println("a2 size:", len(a2))
	fmt.Println("a3 size:", len(a3))
}
