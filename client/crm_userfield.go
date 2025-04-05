package client

func (c *Client) CrmUserfieldFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.userfield.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmUserfieldTypes(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.userfield.types", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmUserfieldEnumerationFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.userfield.enumeration.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmUserfieldSettingsFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.userfield.settings.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
