package main

// #include <stdlib.h>
import "C"
import (
	"fmt"
	"runtime"
)

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	fmt.Printf("%sddd", runtime.GOROOT())
	fmt.Print(Random())
}
