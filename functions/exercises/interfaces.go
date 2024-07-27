package exercises

import (
	"fmt"
	"log"
	"strconv"
)

type Book struct {
	Title string
}

func (b Book) String() string {
	return fmt.Sprint("the titel of the book is ", b.Title)
}

type Count int

func (c Count) String() string {
	return fmt.Sprint("the number: ", strconv.Itoa(int(c)))
}

func LogInfo(s fmt.Stringer) {
	log.Println("log from notes.go", s.String())
}
