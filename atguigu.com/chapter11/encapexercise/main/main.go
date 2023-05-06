package main

import (
	"bufio"
	"bytes"
	"fmt"
	"gocode/atguigu.com/chapter11/encapexercise/model"
	"io"
	"os"
)

func main() {
	//创建一个account变量
	account := model.NewAccount("jzh11111", "000", 40)
	if account != nil {
		fmt.Println("创建成功=", account)
	} else {
		fmt.Println("创建失败")
	}

	var r io.Reader
	r = os.Stdin // see 12.1
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)
	f, _ := os.Open("test.txt")
	r = bufio.NewReader(f)
}
