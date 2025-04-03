package activity_stream

import "github.com/sviridoves/go-bitrix/types"

type ActivityStreamResponse struct {
	types.Response
	Result []ActivityStream `json:"result"`
}
