package exercises

import "fmt"

func VariadicParams(ii ...int) {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)
}

func Foo() {
	fmt.Println("i am foo")
}
func Bar() {
	fmt.Println("i am bar")
}
func Deferring(deferFunc bool) {
	if deferFunc {
		defer Foo()
		Bar()
	} else {
		Foo()
		Bar()
	}
}

type Person struct {
	first string
}

type SecretAgent struct {
	Person
	ltk bool
}

func (p Person) speak() {
	fmt.Println("Hi, I'm ", p.first)
}

func (p SecretAgent) speak() {
	fmt.Println("Hi, I'm secret agent ", p.first)
}

func UseMethod() {
	p := Person{
		first: "Ron",
	}

	p.speak()
}

func WithoutInterface() {
	p := Person{
		first: "Ron",
	}

	sa := SecretAgent{
		Person: p,
		ltk:    true,
	}

	p.speak()
	sa.speak()
}

type human interface {
	speak()
}

func talk(h human) {
	h.speak()
}
func UseInterface() {
	p := Person{
		first: "Ron",
	}

	sa := SecretAgent{
		Person: p,
		ltk:    true,
	}
	talk(p)
	talk(sa)

}
