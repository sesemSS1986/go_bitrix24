package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) ForumMessageUserGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("forum.message.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ForumMessageDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("forum.message.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
