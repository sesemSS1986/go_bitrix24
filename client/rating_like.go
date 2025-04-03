package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) LikeList(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("like.list", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LikeReactions(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("like.reactions", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
