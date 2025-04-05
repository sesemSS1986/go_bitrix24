package client

func (c *Client) CrmWebformList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.webform.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmWebformResultAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.webform.result.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmWebformConfigurationGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.webform.configuration.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
