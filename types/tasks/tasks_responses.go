package tasks

import "github.com/sviridoves/go-bitrix/types"

type TasksResponse struct {
	types.Response
	Result AllTasks `json:"result"`
}
