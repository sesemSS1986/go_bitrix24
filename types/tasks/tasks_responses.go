package tasks

import "github.com/sesemSS1986/go_bitrix24/types"

type TasksResponse struct {
	types.Response
	Result AllTasks `json:"result"`
}
