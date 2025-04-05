package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) Sonet_groupGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupCreate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group.create", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupSetowner(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group.setowner", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
