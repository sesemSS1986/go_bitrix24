package recent_chat_messages

import "github.com/sesemSS1986/go_bitrix24/types"

type MessageListResponse struct {
	types.ResponseTime `json:"time"`
	Result             MessageList `json:"result"`
}
