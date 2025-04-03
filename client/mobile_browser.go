package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) MobileBrowserConstGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("mobile.browser.const.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
