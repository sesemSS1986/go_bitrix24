package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmCatalogFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.catalog.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCatalogGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.catalog.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCatalogList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.catalog.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
