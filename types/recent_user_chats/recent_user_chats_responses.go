package recent_user_chats

import "github.com/sesemSS1986/go_bitrix24/types"

type RecentUserChatsResponce struct {
	types.Response
	Result []RecentUserChats `json:"result"`
}
