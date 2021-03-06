package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//组织一个udp地址结构，指定服务器的IP+port
	//srvAddr是地址结构
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8006")
	//访问命令：nc -u 127.0.0.1 8006
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	fmt.Println("udp 服务器地址结构，创建完成！！！")
	//创建用户通信的socket：udpConn
	udpConn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		fmt.Println("ListenUD err:", err)
		return
	}
	defer udpConn.Close()
	fmt.Println("udp 服务器socket，创建完成！！！")
	//读取客户端发送的数据
	buf := make([]byte, 4096)
	//返回3个值。分别是读取到的字节数，客户端的地址，error
	n, cltAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("ReadFromUDP err:", err)
		return
	}
	//模拟处理数据
	fmt.Printf("服务器读到%vd的数据:%s\n", cltAddr, string(buf[:n]))
	//提取系统当前时间
	daytime := time.Now().String()
	//回写数据给客户端
	_, err = udpConn.WriteToUDP([]byte(daytime), cltAddr)
	if err != nil {
		fmt.Println("WriteToUDP err:", err)
		return
	}

}
