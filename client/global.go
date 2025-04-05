package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) Batch(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("batch", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Scope(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("scope", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Events(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("events", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Profile(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("profile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
