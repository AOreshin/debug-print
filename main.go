package main

import (
	"fmt"
	"time"
)

func main() {
	printLines()
	for i := 0; i < 6; i++ {
		fmt.Printf("\033[A")
		time.Sleep(time.Second)
		fmt.Printf("\033[2K")
		time.Sleep(time.Second)
	}
	printLines()
	time.Sleep(5 * time.Second)
}

func printLines() {
	fmt.Print(`a
b
c
d
e
f
`)
}
