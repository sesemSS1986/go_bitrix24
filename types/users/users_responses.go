package users

import "github.com/sviridoves/go-bitrix/types"

type UsersResponse struct {
	types.Response
	Result []User `json:"result"`
}
