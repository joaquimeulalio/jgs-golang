package main

import (
	"fmt"
)

// Base class
type Base struct {
	valorbase int
}

// FuncPrint on Base class
func (b Base) FuncPrint() {
	// avoid "Print" name clashes ...
	fmt.Printf("Base FuncPrint = %+v\n", b)
}

// Derived1 class
type Derived1 struct {
	//
	// Truque 1 - https://flaviocopes.com/golang-is-go-object-oriented/
	// When declaring a struct we can add an unnamed (anonymous) field, which causes its fields and its methods to be exposed on the struct.
	// This is called struct embedding:
	Base
}

// FuncPrint on Derived1 class
func (b Derived1) FuncPrint() {
	// avoid "Print" name clashes ...
	fmt.Printf("Derived1 FuncPrint = %+v\n", b)
}

// Struct types plus their associated methods serve the same goal of a traditional class, where the struct only holds the state, not the behavior, and the methods provide them behavior, by allowing to change the state.
//  Since methods can only be attached in the same package where the type is defined, we cannot “enrich” the built-in basic types, but we can enrich any named type we create with a base type underlying representation
// Interfaces elegantly provide polymorphism: by accepting an interface, you declare to accept any kind of object satisfying that interface.

// BaseInterface Aqui eh a base do polimorfismo, note que nao precisa ter "type embedded" em contraste com inheritance
// Truque 2
type BaseInterface interface {
	FuncPrint()
}

func main() {

	fmt.Printf("============= C++ class \n")
	b1 := Base{7}
	b1.FuncPrint()
	d1 := Derived1{}
	d1.valorbase = 8
	d1.FuncPrint()
	fmt.Printf("b1=%+v\n", b1)
	fmt.Printf("d1=%+v\n", d1)

	fmt.Printf("\n============= Polimorphism with reference: \n")
	var i BaseInterface
	i = b1
	i.FuncPrint()
	i = d1
	i.FuncPrint()

	// Pode ser com pointers tambem (em C++ o polimorfismo acima seria com "references")
	// Hum, nao funcionou. Mas nao importa, usando o modo acima com references esta otimo.
	/*
		fmt.Printf("============= Polimorphism with pointer: \n")
		var p *BaseInterface
		p = &b1
		p.FuncPrint()
		p = &d1
		p.FuncPrint()
	*/

	fmt.Println("done 8")
}
