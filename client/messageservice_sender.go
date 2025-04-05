package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) MessageserviceSenderAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("messageservice.sender.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MessageserviceSenderDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("messageservice.sender.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MessageserviceSenderList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("messageservice.sender.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
