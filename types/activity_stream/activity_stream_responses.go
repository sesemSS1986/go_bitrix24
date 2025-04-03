package activity_stream

import "github.com/sesemSS1986/go_bitrix24/types"

type ActivityStreamResponse struct {
	types.Response
	Result []ActivityStream `json:"result"`
}
