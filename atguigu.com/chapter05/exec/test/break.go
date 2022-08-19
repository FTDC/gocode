package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("sum:", sum)
			break
		}
	}
}
