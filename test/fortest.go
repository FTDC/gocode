package main

import "fmt"

func main() {
	// for i := 0; i < 5; i++ {
	// 	var v int
	// 	fmt.Printf("%d ", v)
	// 	v = 5
	// }

	// for i := 0; ; i++ {
	// 	fmt.Println("Value of i is now:", i)
	// }

	// for i := 0; i < 3; {
	// 	fmt.Println("Value of i:", i)
	// }

	// s := ""
	// for s != "aaaaa" {
	// 	fmt.Println("Value of s:", s)
	// 	s = s + "a"
	// }

	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
}
