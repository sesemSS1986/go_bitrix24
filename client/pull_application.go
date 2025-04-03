package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) PullApplicationConfigGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("pull.application.config.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) PullApplicationEventAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("pull.application.event.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) PullApplicationPushAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("pull.application.push.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
