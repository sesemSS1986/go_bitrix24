package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmPersontypeFields(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.persontype.fields", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CrmPersontypeList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("crm.persontype.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
