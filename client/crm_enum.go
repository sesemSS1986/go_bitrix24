package client

func (c *Client) CrmEnumSettingsMode(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.enum.settings.mode", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.enum.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumOwnertype(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.enum.ownertype", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumAddresstype(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.enum.addresstype", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumContenttype(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.enum.contenttype", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumActivitytype(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.enum.activitytype", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumActivitypriority(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.enum.activitypriority", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumActivitydirection(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.enum.activitydirection", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumActivitynotifytype(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.enum.activitynotifytype", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumActivitystatus(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.enum.activitystatus", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmEnumEntityeditorConfigurationScope(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.enum.entityeditor.configuration.scope", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
