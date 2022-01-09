package processingcenterdistribution

import (
	"encoding/json"
	"fmt"
	"group-chat/messageType"
	"group-chat/server/model"
	"group-chat/server/tools"
	"net"
)

type UserMessageProcessingStruct struct {
	NetConn net.Conn
	UserId  int
}

func (UserMessageProcessingStruct *UserMessageProcessingStruct) UserMessageProcessing(mes *messageType.MessageStruct) (err error) {
	loginMes := messageType.LoginMessageDataStruct{}
	err = json.Unmarshal([]byte(mes.MessageData), &loginMes)
	if err != nil {
		fmt.Printf("处理登录信息反序列化存入登录消息体内容结构体失败%v\n", err)
		return
	}
	userDaoStruct := model.UserDaoStruct{
		RedisClientPool: model.RedisClientPool,
	}
	loginResMes := messageType.LoginResMessageDataStruct{}
	if user, err := userDaoStruct.LoginVerification(loginMes.UserId, loginMes.UserPwd); err != nil {
		if err == model.ERROR_USER_NOT_EXIST {
			loginResMes.Code = 500 //账号密码错误
			loginResMes.Error = model.ERROR_USER_NOT_EXIST.Error()
			fmt.Printf("服务端验证客户失败、未注册%v\n", err)
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 400 //账号密码错误
			loginResMes.Error = model.ERROR_USER_PWD.Error()
			fmt.Printf("服务端验证客户失败、密码错误%v\n", err)
		} else {
			loginResMes.Code = 505 //账号密码错误
			loginResMes.Error = err.Error()
			fmt.Printf("服务端验证客户失败、未知错误%v\n", err)
		}
	} else {
		loginResMes.Code = 200 //200表示通过
		loginResMes.UserId = UserMessageProcessingStruct.UserId
		loginResMes.UserName = user.UserName
		fmt.Printf("服务端验证用户%v登陆成功\n", user.UserName)
	}

	loginResMesByte, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Printf("登录返回消息结构体序列化失败%v\n", err)
		return
	}
	mes.MessageType = messageType.LoginResMessageType
	mes.MessageData = string(loginResMesByte)

	toolsSerializationProcessingStruct := tools.SerializationProcessingStruct{
		NetConn: UserMessageProcessingStruct.NetConn,
	}
	err = toolsSerializationProcessingStruct.SendMessagePacket(*mes)
	if err != nil {
		fmt.Printf("将要返回给客户端消息发送失败%v\n", err)
		return err
	}

	return
}
