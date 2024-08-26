package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Pointer Exercises")

	// Exercise74()
	// Exercise75()
	// Exercise76()
	Exercise77()
}

func Exercise74() {
	intVal := 42
	intPtr := &intVal

	fmt.Println("pointer address: ", intPtr)
}

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func Exercise75() {
	fmt.Printf("a |value: %v| type %T| data: %v \n", a, a, *a)
	fmt.Printf("b |value: %v| type %T| data: %v \n", b, b, *b)
	fmt.Printf("c |value: %v| type %T| data: %v \n", c, c, *c)
	fmt.Printf("d |value: %v| type %T| data: %v \n", d, d, *d)
}

type dog struct {
	first string
}

func (d *dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

func (d *dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking Woof Woof.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

func Exercise76() {
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()
	// youngRun(d1) // doesn't run

	d2 := &dog{"Padget"}
	d2.walk()
	d2.run()
	youngRun(d2)

}

type person struct {
	first string
}

func changeFirstV(p person, val string) person {
	p.first = val
	return p
}

func changeFirstP(p *person, val string) {
	p.first = val
}

func printName(p person) {
	fmt.Println("My name is ", p.first)
}
func Exercise77() {
	ron := person{
		first: "ron",
	}

	printName(ron)

	ron = changeFirstV(ron, "louis")
	printName(ron)

	changeFirstP(&ron, "sheldon")
	printName(ron)
}
