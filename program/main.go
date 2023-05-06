package main

import "fmt"

func main() {
	//printName("xjin", "张华")
	fr10(11)
}

func fr10(n int) (res int) {
	if n < 0 {
		return
	} else {
		n := n - 1
		fmt.Println(fr10(n))
	}
	return n

}
