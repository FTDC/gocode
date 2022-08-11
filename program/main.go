package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a int64 = 15
	var b bool = false
	// a := 15
	// b := false

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("nihao")
	fmt.Println("nihao")
	fmt.Println("时间的问题")

	fmt.Printf("a 的类型 %T， a 的字节大小 %d \n", a, unsafe.Sizeof(a))

	var (
		ac   int    = 4
		name string = "haha"
	)

	fmt.Printf("a:%d, name:%s", ac, name)
}
