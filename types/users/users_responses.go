package users

import "github.com/sesemSS1986/go_bitrix24/types"

type UsersResponse struct {
	types.Response
	Result []User `json:"result"`
}
