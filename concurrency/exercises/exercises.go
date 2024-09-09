package exercises

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func Ex1() {
	fmt.Println("main started")

	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			fmt.Println("goroutine: ", i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("main finished")
}

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("Hello, my name is", p.name, "and I am", p.age, "years old.")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func Ex2() {
	// 	create a type person struct
	// ● attach a method speak to type person using a pointer receiver
	// ○ *person
	// ● create a type human interface
	// ○ to implicitly implement the interface, a human must have the speak method
	// ● create func “saySomething”
	// ○ have it take in a human as a parameter
	// ○ have it call the speak method
	// ● show the following in your code
	// ○ you CAN pass a value of type *person into saySomething
	// ○ you CANNOT pass a value of type person into saySomething

	randy := person{
		name: "Randy",
		age:  25,
	}
	// randy.speak()
	// speak() only works with ptr to person because speak() is a method with a ptr receiver
	//passing a value of type person into saySomething will not work
	//bc interface speak() method is implemented with a ptr receiver
	saySomething(&randy)
	// saySomething(randy)
}

func Ex3() {
	// Using goroutines, create an incrementer program
	// 	- have a variable to hold the incrementer value
	// 	- launch a bunch of goroutines
	// 		- each goroutine should
	// 			- read the incrementer value
	// 				- store it in a new variable
	// 			- yield the processor with runtime.Gosched()
	// 			- increment the new variable
	// 			- write the value in the new variable back to the incrementer variable
	// 	- use waitgroups to wait for all of your goroutines to finish
	// 	- the above will create a

	maxGoRoutines := 100
	counter := 0
	var wg sync.WaitGroup

	wg.Add(maxGoRoutines)

	for i := 0; i < maxGoRoutines; i++ {
		go func() {
			temp := counter
			runtime.Gosched()
			temp++
			counter = temp
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

func Ex4() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	counter := 0
	maxGoRoutines := 100
	wg.Add(maxGoRoutines)

	for i := 0; i < maxGoRoutines; i++ {
		go func() {
			mu.Lock()
			temp := counter
			temp++
			counter = temp
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}

func Ex5() {

	var wg sync.WaitGroup

	var counter int64 = 0
	maxGoRoutines := 100
	wg.Add(maxGoRoutines)

	for i := 0; i < maxGoRoutines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
