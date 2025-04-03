package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ImopenlinesSessionIntercept(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("imopenlines.session.intercept", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
