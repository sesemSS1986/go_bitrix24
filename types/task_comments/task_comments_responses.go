package task_comments

import "github.com/sviridoves/go-bitrix/types"

type TaskCommentsResponse struct {
	Result             []TaskComments `json:"result"`
	types.ResponseTime `json:"time"`
}
