package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) Sonet_group_subjectGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group_subject.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_group_subjectAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group_subject.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_group_subjectUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group_subject.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_group_subjectDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group_subject.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
