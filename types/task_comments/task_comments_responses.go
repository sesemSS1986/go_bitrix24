package task_comments

import "github.com/sesemSS1986/go_bitrix24/types"

type TaskCommentsResponse struct {
	Result             []TaskComments `json:"result"`
	types.ResponseTime `json:"time"`
}
