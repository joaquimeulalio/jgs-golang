package main

import (
	"errors"
	"fmt"
)

// =============================================================
// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return "=== ERROR STRING ==== " + e.s
}

// New returns an error that formats as the given text.

// =============================================================

// Foo function
func Foo() (string, error) {
	//	return "foo", nil
	//	return "foo", fmt.Errorf("erro que vira instancia de Error()")
	//	return "foo", errorString{"instancia direto"} // nao buildou
	return "foo", errors.New("aqui funciona sem minha propria funcao New()")
}

// Foo1 function
func Foo1() (string, error) {
	return "foo1", nil
}

// Foo2 function
func Foo2() (string, error) {
	return "foo2", fmt.Errorf("erro que vira instancia de Error()")
}

// Foo3 function
func Foo3() (string, error) {
	return "foo3", errors.New("aqui funciona sem minha propria funcao New()")
}

func main() {
	foo, err := Foo()
	if err != nil {
		fmt.Printf("Foo error: %v\n", err)
	}
	fmt.Println(foo)

	foo1, err1 := Foo1()
	if err1 != nil {
		fmt.Printf("Foo1 error: %v\n", err1)
	}
	fmt.Println(foo1)

	foo2, err2 := Foo2()
	if err2 != nil {
		fmt.Printf("Foo2 error: %v\n", err2)
	}
	fmt.Println(foo2)

	foo3, err3 := Foo3()
	if err3 != nil {
		fmt.Printf("Foo3 error: %v\n", err3)
	}
	fmt.Println(foo3)
}
