package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) TaskCtasklogitemGetmanifest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctasklogitem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtasklogitemList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctasklogitem.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
