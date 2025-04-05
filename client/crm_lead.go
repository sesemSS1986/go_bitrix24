package client

func (c *Client) CrmLeadFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadProductrowsSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.productrows.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadProductrowsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.productrows.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadUserfieldAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadUserfieldGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadUserfieldList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.userfield.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadUserfieldUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadUserfieldDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadDetailsConfigurationGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.details.configuration.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadDetailsConfigurationSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.details.configuration.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadDetailsConfigurationReset(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.details.configuration.reset", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmLeadDetailsConfigurationForcecommonscopeforall(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.lead.details.configuration.forcecommonscopeforall", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
