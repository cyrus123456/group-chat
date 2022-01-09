package tools

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"group-chat/messageType"
	"net"
)

type SerializationProcessingStruct struct {
	NetConn net.Conn
}

func (SerializationProcessingStruct *SerializationProcessingStruct) ReceiveMessagePacket() (mes messageType.MessageStruct, err error) {
	Mes_TypeByte := [8096]byte{}
	_, err = SerializationProcessingStruct.NetConn.Read(Mes_TypeByte[:4])
	if err != nil {
		fmt.Printf("服务器读取客户端发送来的信息长度mes_TypeByte失败%v\n", err)
		return
	}
	Mes_TypeByte_Len := binary.BigEndian.Uint32(Mes_TypeByte[:4])
	n, err := SerializationProcessingStruct.NetConn.Read(Mes_TypeByte[:Mes_TypeByte_Len])
	if n != int(Mes_TypeByte_Len) || err != nil {
		fmt.Printf("服务器读取客户端发送来的信息数据传输丢失%v\n", err)
		return
	}

	err = json.Unmarshal(Mes_TypeByte[:Mes_TypeByte_Len], &mes)
	if err != nil {
		fmt.Printf("服务器读取客户端发送来的信息反序列化失败%v\n", err)
		return
	}
	return
}

func (SerializationProcessingStruct *SerializationProcessingStruct) SendMessagePacket(mes messageType.MessageStruct) (err error) {
	mes_jsonByte, err := json.Marshal(mes)
	if err != nil {
		fmt.Printf("服务器返回消息序列化失败%v\n", err)
		return
	}
	Mes_TypeByte := [8096]byte{}
	binary.BigEndian.PutUint32(Mes_TypeByte[:4], uint32(len(mes_jsonByte)))
	n, err := SerializationProcessingStruct.NetConn.Write(Mes_TypeByte[:4])
	if n != 4 || err != nil {
		fmt.Printf("服务器返回消息长度写入发送失败%v\n", err)
		return
	}
	n, err = SerializationProcessingStruct.NetConn.Write(mes_jsonByte)
	if n != len(mes_jsonByte) || err != nil {
		fmt.Printf("服务器返回消息内容失败%v\n", err)
		return
	}
	return
}
