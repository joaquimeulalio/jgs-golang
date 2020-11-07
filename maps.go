// https://stackoverflow.com/questions/40680981/are-maps-passed-by-value-or-by-reference-in-go
package main

import "fmt"

func mapToAnotherFunction(m map[string]int) {
	m["hello"] = 3
	m["world"] = 4
	m["new_word"] = 5
}

// func mapToAnotherFunctionAsRef(m *map[string]int) {
// m["hello"] = 30
// m["world"] = 40
// m["2ndFunction"] = 5
// }

func main() {
	m := make(map[string]int)
	m["hello"] = 1
	m["world"] = 2

	// Initial State
	for key, val := range m {
		fmt.Println(key, "=>", val)
	}

	fmt.Println("-----------------------")

	mapToAnotherFunction(m)
	// After Passing to the function as a pointer
	for key, val := range m {
		fmt.Println(key, "=>", val)
	}

	// Try Un Commenting This Line
	fmt.Println("-----------------------")

	// mapToAnotherFunctionAsRef(&m)
	// // After Passing to the function as a pointer
	// for key, val := range m {
	//  fmt.Println(key, "=>", val)
	// }

	// Outputs
	// hello => 1
	// world => 2
	// -----------------------
	// hello => 3
	// world => 4
	// new_word => 5
	// -----------------------

}
