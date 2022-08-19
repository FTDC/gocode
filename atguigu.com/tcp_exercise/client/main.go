package main

import (
	"bufio"
	"fmt"
	"gocode/atguigu.com/tcp_exercise/config"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {

	// 连接客户端
	conn, err := net.Dial("udp", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		fmt.Errorf("客户端连接失败，错误：%v", err)
		return
	}

	// 关闭连接
	defer conn.Close()

	// 接收输入的内容
	inputRead := bufio.NewReader(os.Stdin)

	for true {
		readString, err := inputRead.ReadString('\n')
		if err != nil {
			fmt.Errorf("接收输入内容失败，错误：%v", err)
			return
		}

		n, err := conn.Write([]byte(readString))
		if err != nil {
			fmt.Errorf("发送数据失败，错误：%v", err)
			return
		}

		log.Print("发送内容", n)

	}

}
