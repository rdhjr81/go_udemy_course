package learning

import (
	"fmt"
	"runtime"
	"sync"
)

func SinglyThreaded() {
	fmt.Println("OS: \t", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	foo()
	bar()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

func UseGoRoutine() {
	fmt.Println("OS: \t", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	go foo()
	go bar()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}

var wg sync.WaitGroup

func UseWaitGroup() {
	fmt.Println("OS: \t", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	bar()
	wg.Wait()

}
