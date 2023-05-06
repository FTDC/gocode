package main

import "fmt" // Package implementing formatted I/O.

func main() {
	var a, b, c int
	a, b, c = 5, 7, "abc"

	fmt.Print(a, b, c)

	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
}
