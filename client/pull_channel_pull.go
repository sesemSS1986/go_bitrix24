package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) PullConfigGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("pull.config.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) PullChannelPublicGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("pull.channel.public.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) PullChannelPublicList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("pull.channel.public.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
