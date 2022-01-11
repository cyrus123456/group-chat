package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"group-chat/messageType"
	"strconv"

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

func (_this *UserDaoStruct) Register(registeredMessageData *messageType.RegisteredMessageDataStruct) (err error) {
	user, err := _this.getUserByid(registeredMessageData.UserId)
	if (registeredMessageData.UserId == user.UserId) && (registeredMessageData.UserPwd == user.UserPwd) {
		fmt.Printf("从Redis中用户已经存在%v\n", err)
		ERROR_USER_EXIST.Error()
		err = ERROR_USER_EXIST
		return
	}
	registeredMessageDataByte, err := json.Marshal(registeredMessageData)
	if err != nil {
		fmt.Printf("注册信息序列化失败%v\n", err)
		return err
	}

	ctx := context.Background()
	RedisClientPool.HSet(ctx, "user", strconv.Itoa(registeredMessageData.UserId), string(registeredMessageDataByte))

	return
}

func (_this *UserDaoStruct) getUserByid(id int) (user *UserStruct, err error) {

	ctx := context.Background()
	userByte, err := _this.RedisClientPool.HGet(ctx, "user", strconv.Itoa(id)).Result()
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
