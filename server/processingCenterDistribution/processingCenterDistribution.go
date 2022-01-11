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

func (_this *ClassificationProcessingStruct) ClassificationProcessing() error {
	for {
		toolsSerializationProcessingStruct := tools.SerializationProcessingStruct{
			NetConn: _this.NetConn,
		}
		mes, err := toolsSerializationProcessingStruct.ReceiveMessagePacket()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("客户端退出%v\n", err)
				return err
			}
			fmt.Printf("服务端读取失败%v\n", err)
			return err
		}
		err = _this.ClassificationJudgment(&mes)
		if err != nil {
			fmt.Printf("服务器信息结构体分类处理失败%v\n", err)
			return err
		}
	}
}

func (_this *ClassificationProcessingStruct) ClassificationJudgment(mes *messageType.MessageStruct) (err error) {
	for {
		switch mes.MessageType {
		case messageType.LoginMessageType:
			fmt.Println("处理登陆消息")
			userMessageProcessingStruct := &UserMessageProcessingStruct{
				NetConn: _this.NetConn,
			}
			err = userMessageProcessingStruct.LoginMessageProcessing(mes)
			if err != nil {
				fmt.Printf("处理登陆消息失败%v\n", err)
				return
			}
			return
		case messageType.RegisteredMessageType:
			fmt.Println("处理注册消息")
			userMessageProcessingStruct := &UserMessageProcessingStruct{
				NetConn: _this.NetConn,
			}
			err = userMessageProcessingStruct.RegistrationMessageProcessing(mes)
			if err != nil {
				fmt.Printf("处理注册消息失败%v\n", err)
				return
			}
			return
		default:
			fmt.Println("消息类型错误")
			return
		}
	}
}
