package learning

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func (p Person) WriteOut(w io.Writer) error {
	_, err := w.Write([]byte(p.First))
	if err != nil {
		log.Fatalf("error %s", err)
	}
	return err
}

func UseBuffer() {
	b := bytes.NewBufferString("Hello ")
	fmt.Println(b.String())

	b.WriteString(" Rob")
	fmt.Println(b.String())

	b.Reset()
	b.WriteString("Banana!")
	fmt.Println(b.String())

}

func WriteToFile() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	p := Person{
		First: "Ron Don",
	}

	var b bytes.Buffer

	p.WriteOut(f)
	p.WriteOut(&b)

	fmt.Println(b)

	b2 := []byte{82, 111, 110, 42, 69, 111, 111}
	b2String := string(b2)
	fmt.Println(b2String)
}

func foo() {
	fmt.Println("foo")
}
func AnonymousFunctions() {

	//no arg
	func() {
		fmt.Println("we are anoymous")
	}()

	//arg
	func(s string) {
		fmt.Println("hello ", s)
	}("Arthur Pendragon")
}

func FunctionExpression() {

	//no arg
	x := func() {
		fmt.Println("we are anoymous")
	}

	x()
	x()

	//arg
	y := func(s string) {
		fmt.Println("hello ", s)
	}

	y("george")
	y("alan")
}

func ReturnAFunction() {
	x := foo2()
	fmt.Println(x)

	y := bar()
	result := y()
	fmt.Println(result)

	fmt.Printf("%T bar:	", bar)
	fmt.Println()
	fmt.Printf("%T y: ", y)
}

func foo2() int {
	return 42
}

func bar() func() int {
	return foo2
}

func functionCallback(fInt func() int) {
	fmt.Println("fInt: ", fInt())
}

func UseFunctionAsCallback() {
	functionCallback(bar())
}
