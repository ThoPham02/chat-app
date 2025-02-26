// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.6

package types

type Conversation struct {
	ConversationID int64  `json:"conversationID"`
	Conversation   string `json:"conversation"`
	AvatarUrl      string `json:"avatarUrl"`
	LastMessage    string `json:"lastMessage"`
	LastTime       int64  `json:"lastTime"`
	UnreadCount    int64  `json:"unreadCount"`
}

type LoginReq struct {
	UserName string `json:"userName"` // user_name
	Password string `json:"password"` //  password
}

type LoginRes struct {
	Result Result `json:"result"`
	Token  string `json:"token"` // jwt token for api
	User   User   `json:"user"`  // Account info
}

type Message struct {
	MessageID      int64  `json:"messageID"`
	ConversationID int64  `json:"conversationID"`
	SenderID       int64  `json:"senderID"`
	Message        string `json:"message"`
	CreatedAt      int64  `json:"createdAt"`
}

type RegisterReq struct {
	UserName string `json:"userName"` // user_name
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

type RegisterRes struct {
	Result Result `json:"result"`
	Token  string `json:"token"` // jwt token for api
	User   User   `json:"user"`  // Account info
}

type Result struct {
	Code    int    `json:"code"`    //    Result code: 0 is success. Otherwise, getting an error
	Message string `json:"message"` // Result message: detail response code
}

type User struct {
	UserID    int64  `json:"userID"`
	FullName  string `json:"fullName"`
	AvatarUrl string `json:"avatarUrl"`
	CreatedAt int64  `json:"createdAt"`
}
