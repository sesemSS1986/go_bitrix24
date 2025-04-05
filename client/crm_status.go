package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmStatusFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.status.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.status.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.status.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.status.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.status.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.status.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusEntityTypes(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.status.entity.types", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusEntityItems(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.status.entity.items", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusExtraFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.status.extra.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
