package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ListsGetIblockTypeId(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("lists.get.iblock.type.id", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
