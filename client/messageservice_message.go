package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) MessageserviceMessageStatusUpdate(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("messageservice.message.status.update", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
