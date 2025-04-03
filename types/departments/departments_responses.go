package departments

import "github.com/sviridoves/go-bitrix/types"

type DepartmentResponse struct {
	types.Response
	Result []Department `json:"result"`
}
