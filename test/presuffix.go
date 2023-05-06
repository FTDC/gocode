package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?  \n", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))

	var str2 string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	fmt.Printf("Number of H's in %s is: ", str2)
	fmt.Printf("%d\n", strings.Count(str2, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))

	r := strings.NewReader(str2)
	fmt.Println("r=: %p", r)

}
