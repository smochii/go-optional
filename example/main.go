package main

import (
	"fmt"

	"github.com/smochii/go-optional"
)

type S struct {
	A optional.Optional[string]
	B optional.Optional[int]
	C optional.Optional[float64]
}

func main() {
	s := S{
		A: optional.New("hello"),
		B: optional.New(123),
	}

	// This will be printed
	if s.A.IsPresent() {
		fmt.Println(s.A.Get())
	}

	// This will be printed
	if s.B.IsPresent() {
		fmt.Println(s.B.Get())
	}

	// This will not be printed
	if s.C.IsPresent() {
		fmt.Println(s.C.Get())
	}

	// This will be printed
	fmt.Println(s.C.OrElse(123.456))
}
