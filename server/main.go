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
}
