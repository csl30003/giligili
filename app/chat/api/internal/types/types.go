// Code generated by goctl. DO NOT EDIT.
package types

type ChatMessage struct {
	Id               int64  `json:"id"`
	FromUserId       int64  `json:"fromUserId"`
	FromUserNickname string `json:"fromUserNickname"`
	ToUserId         int64  `json:"toUserId"`
	ToUserNickname   string `json:"toUserNickname"`
	Content          string `json:"content"`
	CreateTime       string `json:"createTime"`
}

type SendChatMessageReq struct {
	ToUserId int64  `json:"toUserId"`
	Content  string `json:"content"`
}

type SendChatMessageResp struct {
}

type GetChatHistoryReq struct {
	ToUserId int64 `json:"toUserId"`
}

type GetChatHistoryResp struct {
	ChatHistory []ChatMessage `json:"chatHistory"`
}

type DeleteChatMessageReq struct {
	Id int64 `json:"id"`
}

type DeleteChatMessageResp struct {
	Success bool `json:"success"`
}
