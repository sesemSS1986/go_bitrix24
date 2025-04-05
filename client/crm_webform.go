package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmWebformList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.webform.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmWebformResultAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.webform.result.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmWebformConfigurationGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.webform.configuration.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
