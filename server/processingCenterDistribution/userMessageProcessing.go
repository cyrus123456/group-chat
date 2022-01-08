package processingcenterdistribution

import (
	"encoding/json"
	"fmt"
	"group-chat/messageType"
	"net"
)

type UserMessageProcessingStruct struct {
	NetConn net.Conn
	UserId  string
}

func (UserMessageProcessingStruct *UserMessageProcessingStruct) UserMessageProcessing(mes *messageType.MessageStruct) (err error) {
	loginMes := messageType.LoginMessageDataStruct{}
	err = json.Unmarshal([]byte(mes.MessageData), &loginMes)
	if err != nil {
		fmt.Printf("处理登录信息反序列化存入登录消息体内容结构体失败%v\n", err)
		return
	}
	loginResMes = messageType.LoginResMessageDataStruct{}
	return
}
