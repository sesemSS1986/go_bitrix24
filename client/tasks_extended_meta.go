package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) Tasks_extendedMetaSetanystatus(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("tasks_extended.meta.setanystatus", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Tasks_extendedMetaOccurinlogsas(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("tasks_extended.meta.occurinlogsas", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
