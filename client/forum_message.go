package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ForumMessageUserGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("forum.message.user.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ForumMessageDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("forum.message.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
