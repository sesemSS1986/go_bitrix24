package client

import "github.com/sviridoves/go-bitrix/types"

func (c *Client) CrmDuplicateFindbycomm(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.duplicate.findbycomm", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
