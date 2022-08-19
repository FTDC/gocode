package main

import (
	"fmt"
	"gocode/atguigu.com/tcp_exercise/config"
	"net"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(config.Host + ":" + strconv.Itoa(config.Port))
	listen, err := net.Listen("udp", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		fmt.Errorf("监听端口失败...., 错误：%v", err)
		return
	}
	for {
		// 和客户端建立连接
		conn, err := listen.Accept()
		if err != nil {
			//
			fmt.Errorf("客户端建立连接失败, 错误：%v", err)
			break
		}

		// 协程处理
		go tcpHandShark(conn)
	}

}

func tcpHandShark(conn net.Conn) {

	fmt.Println("创建连接成功 ！！！！", conn.RemoteAddr().String())

	// 关闭连接
	defer conn.Close()

	for {
		// 声明切片的接收内容
		buff := make([]byte, 1024)
		// 读取内容
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Errorf("接收内容失败，错误： %s", err)
		}
		fmt.Println("服务端收到的内容：", strings.Trim(string(buff[0:n]), "\n"))
	}
}
