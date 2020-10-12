package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fmt.Println("File creation")
	//permissions := 0644 // or whatever you need
	//byteArray := []byte("to be written to a file\n")
	//err := ioutil.WriteFile("file.txt", byteArray, permissions)

	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	if err != nil {
		panic(err)
	}

}
