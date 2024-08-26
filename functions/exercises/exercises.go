package exercises

import (
	"fmt"
	"math"
)

// the foo func
func Foo() int {
	return 42
}

// the bar func
func Bar() (int, string) {
	return 43, "fooooobar"
}

func Exercise59() {
	foo := func(ints ...int) int {
		sum := 0
		for _, v := range ints {
			sum += v
		}
		return sum
	}
	ints := []int{1, 2, 3, 4, 5, 6, 7}
	sum := foo(ints...)
	fmt.Println("foo sum: ", sum)

	bar := func(ints []int) int {
		return foo(ints...)
	}
	sum = bar(ints)
	fmt.Println("bar sum: ", sum)
}

func Exercise60() {
	f := func(funcName string) {
		fmt.Println("I am the func with name ", funcName)
	}
	f("func 1")
	defer f("deferred func 1")
	defer f("deferred func 2")
	f("func 2")
}

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("my name is ", p.first, " and I am , ", p.age, " years old.")
}

func Exercise61() {
	alan := person{
		first: "alan",
		age:   42,
	}
	alan.speak()
}

type square struct {
	length float64
	width  float64
}

// function that calcs area of a square
func (sq square) area() float64 {
	return sq.length * sq.width
}

type circle struct {
	width float64
}

func (c circle) area() float64 {
	return math.Pi * c.width
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("area of shape is: ", s.area())
}

func Exercise62() {
	sq := square{
		length: 10,
		width:  4.2,
	}
	info(sq)

	c := circle{
		width: 54.2,
	}

	info(c)
}
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
func add(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b
}

func Exercise68() {

	c := 4
	d := 5
	func() {
		d = c
	}()
	fmt.Println("d = ", d)
}

func Exercise69() {
	add := func(a int, b int) int {
		return a + b
	}
	result := add(1, 2)
	fmt.Println("1 + 2 = ", result)
}

func get42Func() func() int {
	return func() int {
		return 42
	}
}
func Exercise70() {
	funk42 := get42Func()
	fmt.Printf("func 42 type: %t", funk42)
	fmt.Println()
	fmt.Printf("func 42 value: %d", funk42())
}

func Exercise71() {
	func(f func() int) {
		fmt.Printf("callback type: %t", f)
		fmt.Println()
		fmt.Printf("callback value: %d", f())
	}(get42Func())
}

func Exercise72() {
	index := 1
	func() {
		for ; index < 5; index++ {
			fmt.Println("index: ", index)
		}
		fmt.Println()
	}()
}
