package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ImbotMessageAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.message.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotMessageDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.message.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotMessageUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.message.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotMessageLike(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.message.like", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
