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
}
