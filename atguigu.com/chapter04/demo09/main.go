package main

import (
	"fmt"
)

func main() {

	//位运算的演示
	fmt.Println(2 & 3)  // 2
	fmt.Println(2 | 3)  // 3
	fmt.Println(2 ^ 3)  // 3
	fmt.Println(-2 ^ 2) //-4

	a := 1 >> 2 //0
	c := 1 << 2 //4
	fmt.Println("a=", a, "c=", c)

	//  -1 原码 1000 0010  反码： 1111 11101 补码  1111 1110
	d := -2 >> 2 //

	fmt.Println("nem", d)

}
