package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	s1 := puppy.Bark()
	fmt.Println(s1)
	s2 := puppy.Barks()
	fmt.Println(s2)

	s3 := puppy.BigBark()
	fmt.Println(s3)

	s4 := puppy.BigBarks()
	fmt.Println(s4)
}
