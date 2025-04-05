package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ImMessageAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.message.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImMessageDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.message.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImMessageUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.message.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImMessageLike(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.message.like", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImMessageCommand(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.message.command", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImMessageShare(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.message.share", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImMessageUserGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.message.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
