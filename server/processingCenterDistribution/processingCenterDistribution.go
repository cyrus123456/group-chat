package processingcenterdistribution

import (
	"fmt"
	"group-chat/messageType"
	"group-chat/server/tools"
	"io"
	"net"
)

type ClassificationProcessingStruct struct {
	NetConn net.Conn
}

func (ClassificationProcessingStruct *ClassificationProcessingStruct) ClassificationProcessing() (err error) {
	toolsSerializationProcessingStruct := tools.SerializationProcessingStruct{
		NetConn: ClassificationProcessingStruct.NetConn,
	}
	mes, err := toolsSerializationProcessingStruct.ReceiveMessagePacket()
	if err != nil {
		if err == io.EOF {
			fmt.Printf("客户端退出%v\n", err)
			return
		}
		fmt.Printf("服务端读取失败%v\n", err)
		return
	}
	err = ClassificationProcessingStruct.ClassificationJudgment(&mes)
	if err != nil {
		fmt.Printf("服务器信息结构体分类处理失败%v\n", err)
		return err
	}
	return
}

func (ClassificationProcessingStruct *ClassificationProcessingStruct) ClassificationJudgment(mes *messageType.MessageStruct) (err error) {
	for {
		switch mes.MessageType {
		case messageType.LoginMessageType:
			fmt.Println("处理登陆消息")
			return
		default:
			fmt.Println("消息类型错误")
			return
		}
	}
}
