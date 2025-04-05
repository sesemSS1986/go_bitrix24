package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmUserfieldFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.userfield.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmUserfieldTypes(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.userfield.types", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmUserfieldEnumerationFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.userfield.enumeration.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmUserfieldSettingsFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.userfield.settings.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
