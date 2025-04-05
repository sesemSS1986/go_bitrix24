package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CrmExternalchannelConnectorFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.externalchannel.connector.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmExternalchannelConnectorList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.externalchannel.connector.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmExternalchannelConnectorRegister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.externalchannel.connector.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmExternalchannelConnectorUnregister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.externalchannel.connector.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmExternalchannelCompany(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.externalchannel.company", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmExternalchannelContact(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.externalchannel.contact", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmExternalchannelActivityCompany(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.externalchannel.activity.company", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmExternalchannelActivityContact(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.externalchannel.activity.contact", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
