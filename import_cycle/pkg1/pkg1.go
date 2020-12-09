package pkg1

import (
	"fmt"

	"../pkg2"
)

// IntType1 IntType1
type IntType1 int

// TestPkg1 TestPkg1
func TestPkg1() {
	i1 := IntType1(111)
	fmt.Printf("TestPkg1 %d\n", i1)
	pkg2.TestPkg2()
}
