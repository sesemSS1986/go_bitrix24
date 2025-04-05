package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmProductrowFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.productrow.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductrowAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.productrow.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductrowGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.productrow.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductrowList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.productrow.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductrowUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.productrow.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductrowDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.productrow.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
