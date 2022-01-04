package main

import (
	"fmt"
	"net"
)

func main() {
	netListener, error := net.Listen("tcp", "0.0.0.0:8888")
	defer netListener.Close()
	if error != nil {
		fmt.Printf("\"0.0.0.0:8888\"端口监听失败%v\n", error)
	} else {
		fmt.Println("\"0.0.0.0:6666\"端口监听成功")
	}

	// 循环监听
	for {
		netConn, error := netListener.Accept()
		if error != nil {
			fmt.Print("sever main netConn 循环中接收消息失败")
		}
		go distributionLink(netConn)
	}
}

func distributionLink(netConn net.Conn) {
	defer netConn.Close()

	panic("unimplemented")
}
