package main

import (
	"io"
	"log"
)

func func1(s string) (n int, err error) {
	n = 3
	defer func() {
		log.Printf("func2(%q) = %d, %v", s, n, err)
	}()
	n = 23
	return n, io.EOF
}

func main() {
	func1("Go")
}
