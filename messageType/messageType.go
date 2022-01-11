package messageType

//消息类型常量
const (
	LoginMessageType           = "loginMessageType"
	LoginResMessageType        = "loginResMessageType"
	RegisteredMessageType      = "registeredMessageType"
	RegisteredResMessageType   = "registeredResMessageType"
	UserStatusNotificationType = "userStatusNotificationType"
	SmsMessageType             = "smsMessageType"
	SmsResMessageType          = "smsResMessageType"
)

//用户状态常量
const (
	UserIsOnline  = "userIsOnline"
	userIsOffline = "userIsOffline"
	userIsBusy    = "userIsBusy"
)

type MessageStruct struct {
	MessageType string `json:"messageType"`
	MessageData string `json:"messageData"`
}

type LoginMessageDataStruct struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMessageDataStruct struct {
	OnLineUsersId []int  `json:"onLineUsersId"`
	UserId        int    `json:"userId"`
	Code          int    `json:"code"` //返回的状态码
	Error         string `json:"error"`
	UserName      string `json:"userName"`
}

type RegisteredMessageDataStruct struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type RegisteredResMessageDataStruct struct {
	OnLineUsersId []int  `json:"onLineUsersId"`
	UserId        int    `json:"userId"`
	Code          int    `json:"code"` //返回的状态码
	Error         string `json:"error"`
	UserName      string `json:"userName"`
}

type UserStatusNotificationStruct struct {
	UserId            int    `json:"userId"`
	UserCurrentStatus string `json:"userCurrentStatus"`
}

type UserStatusStruct struct {
	UserId            int    `json:"userId"`
	UserPwd           string `json:"userPwd"`
	UserName          string `json:"userName"`
	UserCurrentStatus string `json:"userCurrentStatus"`
}

type SmsMessageDataStruct struct {
	Content    string `json:"content"`
	UserStatus string `json:"userStatus"`
}

type SmsResMessageDataStruct struct {
	Content    string `json:"content"`
	UserStatus string `json:"userStatus"`
}
