package client

func (c *Client) CrmExternalchannelConnectorFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.externalchannel.connector.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmExternalchannelConnectorList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.externalchannel.connector.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmExternalchannelConnectorRegister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.externalchannel.connector.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmExternalchannelConnectorUnregister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.externalchannel.connector.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmExternalchannelCompany(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.externalchannel.company", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmExternalchannelContact(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.externalchannel.contact", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmExternalchannelActivityCompany(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.externalchannel.activity.company", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmExternalchannelActivityContact(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.externalchannel.activity.contact", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
