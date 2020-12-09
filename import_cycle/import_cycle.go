package main

import (
	"fmt"

	"./pkg1"
	"./pkg2"
)

func main() {
	fmt.Println("hello world")
	pkg1.TestPkg1()
	pkg2.TestPkg2()
}
