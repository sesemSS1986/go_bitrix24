package client

import "github.com/sviridoves/go-bitrix/types"

func (c *Client) MobileFormProfile(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.form.profile", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
