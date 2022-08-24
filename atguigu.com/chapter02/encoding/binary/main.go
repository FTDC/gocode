package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	// 将一个字符串转化为一个字节流数组
	var data []byte = []byte("你好")
	fmt.Println("data len", len(data))

	// uint32类型的数字在字节流切片上占4个字节
	// PutUint32()专门用来处理固定长度的数字
	buf := make([]byte, 6+len(data))
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	fmt.Printf("固定长度，%d,%s", buf[:4], buf[4:])

	//copy函数将data数据赋值到buf[4:]中
	copy(buf[4:], data)
	fmt.Println(len(data))
	fmt.Printf("拷贝后%d,%s", buf[:4], buf[4:])

}
