package main

import "fmt"

func main() {
	// var arrAge = [5]int{18, 20, 15, 22, 16}
	var arrKeyValue = [...]int{5, 6, 22}
	// var arrKeyValue = []int{5, 8, 22} //注：初始化得到的实际上是切片slice
	// var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	// var arrKeyValue = []string{3: "Chris", 4: "Ron"}	//注：初始化得到的实际上是切片slice

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %d\n", i, arrKeyValue[i])
	}
}
