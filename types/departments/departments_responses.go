package departments

import "github.com/sesemSS1986/go_bitrix24/types"

type DepartmentResponse struct {
	types.Response
	Result []Department `json:"result"`
}
