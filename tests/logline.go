package main

import (
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func logLine(s string) {
	f, err := os.OpenFile("/tmp/abc.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	f.WriteString(s)
	f.WriteString("\n")
	f.Close()
}

func main() {
	log.Println("============================")
	log.Println("logline 1.13 starting ...")
	logLine("logline 1.13 starting ...")
	logLine("logline 1.13 starting ... line 2")

}
