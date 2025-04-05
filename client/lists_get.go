package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ListsGetIblockTypeId(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.get.iblock.type.id", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
