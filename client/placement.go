package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) PlacementList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("placement.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) PlacementBind(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("placement.bind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) PlacementUnbind(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("placement.unbind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) PlacementGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("placement.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
