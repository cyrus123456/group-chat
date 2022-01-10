package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"group-chat/messageType"

	"github.com/go-redis/redis"
)

var (
	ERROR_USER_NOT_EXIST = errors.New("用户不存在")
	ERROR_USER_EXIST     = errors.New("用户已存在")
	ERROR_USER_PWD       = errors.New("密码错误")
)

type UserStruct struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type UserDaoStruct struct {
	RedisClientPool *redis.Client
}

func (_this *UserDaoStruct) LoginVerification(userId int, userPwd string) (user *UserStruct, err error) {
	if user, err = _this.getUserByid(userId); err == redis.Nil {
		fmt.Printf("从Redis获取用户失败，用户不存在%v\n", err)
		ERROR_USER_NOT_EXIST.Error()
		err = ERROR_USER_NOT_EXIST
		return
	} else if err != nil {
		panic(err)
	}
	if userPwd != user.UserPwd {
		ERROR_USER_PWD.Error()
		err = ERROR_USER_PWD
		return
	}
	return
}

func (_this *UserDaoStruct) register(registeredMessageData *messageType.RegisteredMessageDataStruct) (err error) {
	user, err := _this.getUserByid(registeredMessageData.UserId)
	if (registeredMessageData.UserId == user.UserId) && (registeredMessageData.UserPwd == user.UserPwd) {
		fmt.Printf("从Redis中用户已经存在%v\n", err)
		ERROR_USER_EXIST.Error()
		err = ERROR_USER_EXIST
		return
	}

	if registeredMessageDataByte, err := json.Marshal(registeredMessageData); err != nil {
		fmt.Printf("注册信息序列化失败%v\n", err)
		return err
	}
	return
}

func (_this *UserDaoStruct) getUserByid(id int) (user *UserStruct, err error) {

	ctx := context.Background()
	userByte, err := _this.RedisClientPool.Get(ctx, string(id)).Result()
	if err == redis.Nil {
		fmt.Printf("从Redis获取用户失败，用户不存在%v\n", err)
		ERROR_USER_NOT_EXIST.Error()
		err = ERROR_USER_NOT_EXIST
		return
	} else if err != nil {
		panic(err)
	}

	user = &UserStruct{}
	if err = json.Unmarshal([]byte(userByte), user); err != nil {
		fmt.Printf("从Redis获取用户信息反序列化为可操作的用户结构体失败%v\n", err)
		return
	}

	return

}
