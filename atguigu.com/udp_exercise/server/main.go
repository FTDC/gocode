package main

import (
	"fmt"
	"gocode/atguigu.com/udp_exercise/config"
	"net"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(config.Host + ":" + strconv.Itoa(config.Port))
	conn, err := net.ListenPacket("udp", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		fmt.Errorf("udp监听端口失败...., 错误：%v", err)
		return
	}
	for {
		buff := make([]byte, config.UdpBufSize)
		// 和客户端建立连接
		n, addr, err := conn.ReadFrom(buff)

		if err != nil {
			//
			fmt.Errorf("客户端建立连接失败, 错误：%v", err)
			break
		}

		println("接收的内容:", string(buff[:n]), "addr:", addr.String())
		// 协程处理
		//go udpHandShark(conn)
	}

}

func udpHandShark(conn net.Conn) {

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
