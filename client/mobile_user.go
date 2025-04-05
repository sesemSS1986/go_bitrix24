package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) MobileUserGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileUserCanusetelephony(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.user.canusetelephony", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
