package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) SmileGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("smile.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
