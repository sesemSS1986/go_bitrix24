package recent_user_chats

import "github.com/sviridoves/go-bitrix/types"

type RecentUserChatsResponce struct {
	types.Response
	Result []RecentUserChats `json:"result"`
}
