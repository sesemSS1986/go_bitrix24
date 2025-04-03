package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) Batch(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("batch", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Scope(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("scope", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Events(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("events", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) Profile(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("profile", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
