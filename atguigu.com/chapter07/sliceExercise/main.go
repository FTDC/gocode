package main

import (
	"fmt"
)

func test(slice []int) {
	slice[0] = 100 //这里修改slice[0],会改变实参
}

func main() {

	var slice = []int{1, 2, 3, 4}
	fmt.Println("slice=", slice) // [1,2,3,4]
	test(slice)
	fmt.Println("slice=", slice) // [100, 2, 3, 4]

	ints := enlarge(slice, 2)

	fmt.Println("enlarge size=", cap(ints))

	s := []string{"M", "N", "O", "P", "Q", "R"}
	in := []string{"A", "B", "C"}
	res := InsertStringSlice(s, in, 0)  // at the front
	fmt.Println(res)                    // [A B C M N O P Q R]
	res2 := InsertStringSlice(s, in, 3) // [M N O A B C P Q R]
	fmt.Println(res2)

}

func InsertStringSlice(s []string, in []string, i int) interface{} {
	i2 := len(in)
	strings := make([]string, len(s)+i2)
	copy(strings, s[0:i])
	i3 := append(strings[i:i], in...)
	copy(strings[i+i2:], s[i:])
	fmt.Println("is======", i3)
	//copy(strings, s[i2:])

	return strings
}

func enlarge(arr []int, size int) []int {
	ns := make([]int, cap(arr)*size)

	copy(ns, arr)
	return ns
}
