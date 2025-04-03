package recent_chat_messages

import "github.com/sviridoves/go-bitrix/types"

type MessageListResponse struct {
	types.ResponseTime `json:"time"`
	Result             MessageList `json:"result"`
}
