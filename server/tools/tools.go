package tools

import (
	"fmt"
	"group-chat/messageType"
	"net"
)

type SerializationProcessingStruct struct {
	NetConn net.Conn
}

func (this *SerializationProcessingStruct) ReadMessagePacket() (mes messageType.MessageStruct, Error error) {
	var Mes_TypeByte [8096]byte
	_, err := this.NetConn.Read(Mes_TypeByte[0:4])
	if err != nil {
		fmt.Printf("服务器读取客户端发送来的信息长度mes_TypeByte失败%v\n", err)
		return
	}
	return
}
