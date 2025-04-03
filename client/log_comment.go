package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) LogCommentUserGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("log.comment.user.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) LogCommentDelete(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("log.comment.delete", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
