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
		return
	} else {
		fmt.Println("\"0.0.0.0:6666\"端口监听成功")
	}
	//监听成功持续等待链接，防止退出程序
	for {
		fmt.Println("等待客户端消息")
		netConn, err := netListener.Accept()
		if err != nil {
			fmt.Printf("接收消息失败%v\n", err)
		}

		go processingCenter(netConn)
	}
}

func processingCenter(netConn net.Conn) {
	defer netConn.Close()
}
