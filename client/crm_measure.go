package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmMeasureFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.measure.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmMeasureAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.measure.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmMeasureGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.measure.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmMeasureList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.measure.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmMeasureUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.measure.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmMeasureDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.measure.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
