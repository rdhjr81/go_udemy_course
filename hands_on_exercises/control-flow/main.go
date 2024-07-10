package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

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

func ex27() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	println("x is ", x, ", y is ", y)

	var status string
	if x < 4 && y < 4 {
		status = "x and y are less than 4"
	} else if x > 6 && y > 6 {
		status = "x and y are greater than 6"
	} else if x >= 4 && x <= 6 {
		status = "x is between 4 and 6"
	} else if y != 5 {
		status = "y is not 5"
	} else {
		status = "no conditions met"
	}
	fmt.Println(status)
}

func ex28() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	println("x is ", x, ", y is ", y)

	var status string

	switch {
	case x < 4 && y < 4:
		status = "x and y are less than 4"
	case x > 6 && y > 6:
		status = "x and y are greater than 6"
	case x >= 4 && x <= 6:
		status = "x is between 4 and 6"
	case y != 5:
		status = "y is not 5"
	default:
		status = "no conditions met"
	}

	fmt.Println(status)
}

func ex29(numTimesToPrint int) {
	for i := 0; i < numTimesToPrint; i++ {
		fmt.Println()
		for j := 0; j < 100; j++ {
			fmt.Print(j, ",")
		}
	}
}

func ex30() {
	for i := 0; i < 42; i++ {
		fmt.Println("iteration", i)
		x := rand.Intn(5)
		switch x {
		case 0:
			fmt.Println("x is 0")
		case 1:
			fmt.Println("x is 1")
		case 2:
			fmt.Println("x is 2")
		case 3:
			fmt.Println("x is 3")
		case 4:
			fmt.Println("x is 4")
		default:
			fmt.Println("x is not 0, 1, 2, 3, or 4")
		}
	}
}

func ex31() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func ex32() {
	i := 0
	max := 10
	randomBreak := rand.Intn(max)
	for i < max {
		fmt.Println(i)
		if i == randomBreak {
			fmt.Println("breaking out of loop, random break was ", randomBreak)
			break
		}
		i++
	}
}

func ex33() {
	for i := 0; i < 99; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
}

func ex35() {
	xi := []int{42, 43, 44, 45, 46, 47}
	for i, v := range xi {
		fmt.Println(i, v)
	}
}

func ex36() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	for i, j := range m {
		fmt.Println(i, j)
	}

	if age, ok := m["James"]; ok {
		fmt.Println("James' age is", age)
	}

	qAge := m["Q"]
	fmt.Println("Q's age is", qAge)

	if _, ok := m["Q"]; !ok {
		fmt.Println("Q's age is not in the map")
	}
}

func ex38() {
	num3s := 0
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("found 3")
			num3s++
		}
	}
	fmt.Println("total number of 3s found", num3s)
}

func main() {
	// ex24()
	// ex25()
	// ex27()
	// ex28()
	// ex29(1)
	// ex29(100)
	// ex30()
	// ex31()
	// ex32()
	// ex33()
	// ex35()
	// ex36()
	ex38()
}
