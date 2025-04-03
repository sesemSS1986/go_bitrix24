package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ImopenlinesNetworkJoin(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.network.join", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) ImopenlinesNetworkMessageAdd(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.network.message.add", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
