package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		res := clac(i)
		switch res {
		case 2:
			fmt.Println("Fizz")
		case 3:
			fmt.Println("Buzz")
		case 4:
			fmt.Println("FizzBuzz")
		default:
			fmt.Println(i)

		}
	}
}

// clac  计算数组类型
func clac(num int) (typen int) {
	if (num % 3) == 0 {
		typen = 2
	}

	if (num % 5) == 0 {
		typen = 3
	}

	if (num % 15) == 0 {
		typen = 4
	}

	return typen
}
