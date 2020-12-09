package pkg2

import (
	"fmt"
	// imports ../pkg1: import cycle not allowed
	//	"../pkg1"
)

// IntType2 IntType2
type IntType2 int

// TestPkg2 TestPkg2
func TestPkg2() {
	fmt.Println("TestPkg2")
	//i1 := pkg1.IntType1(111)
}
