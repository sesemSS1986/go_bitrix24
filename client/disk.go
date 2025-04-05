package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) Download(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("download", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Upload(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("upload", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
