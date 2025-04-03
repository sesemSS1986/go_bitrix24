package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) TaskCtasklogitemGetmanifest(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctasklogitem.getmanifest", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) TaskCtasklogitemList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("task.ctasklogitem.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
