package main

import (
	"fmt"
	"math/rand"
)

func ex24() {
	x := rand.Intn(251)
	fmt.Println("value of x is", x)

	if x >= 0 && x <= 100 {
		fmt.Println("between 0 and 100")
	} else if x >= 10 && x <= 200 {
		fmt.Println("between 101 and 200")
	} else {
		fmt.Println("between 201 and 250")
	}
}

func ex25() {
	x := rand.Intn(251)
	fmt.Println("value of x is", x)
	switch {
	case x >= 0 && x <= 100:
		fmt.Println("between 0 and 100")
	case x > 100 && x < 201:
		fmt.Println("between 101 and 200")
	default:
		fmt.Println("between 201 and 250")
	}
}

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	// ex24()
	ex25()
}
