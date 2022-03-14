package main

import (
	"fmt"
	"net"
)

func main() {
	//指定服务器 ip+port创建通信套接字
	conn, err := net.Dial("udp", "127.0.0.1:8006")
	if err != nil {
		fmt.Println("net.dial err:", err)
		return
	}
	defer conn.Close()
	//主动写数据给服务器
	conn.Write([]byte("are you ready?"))
	buf := make([]byte, 4096)
	//接收服务器回发的数据
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	//处理数据，打印
	fmt.Println("服务器回发：", string(buf[:n])) //读了多少显示多少
}
