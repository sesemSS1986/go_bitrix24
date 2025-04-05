package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmEnumSettingsMode(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.enum.settings.mode", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.enum.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumOwnertype(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.enum.ownertype", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumAddresstype(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.enum.addresstype", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumContenttype(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.enum.contenttype", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumActivitytype(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.enum.activitytype", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumActivitypriority(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.enum.activitypriority", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumActivitydirection(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.enum.activitydirection", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumActivitynotifytype(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.enum.activitynotifytype", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumActivitystatus(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.enum.activitystatus", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumEntityeditorConfigurationScope(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.enum.entityeditor.configuration.scope", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
