package client

func (c *Client) CrmCompanyFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyContactFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.contact.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyContactAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.contact.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyContactDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.contact.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyContactItemsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.contact.items.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyContactItemsSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.contact.items.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyContactItemsDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.contact.items.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyUserfieldAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyUserfieldGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyUserfieldList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.userfield.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyUserfieldUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyUserfieldDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyDetailsConfigurationGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.details.configuration.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyDetailsConfigurationSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.details.configuration.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyDetailsConfigurationReset(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.details.configuration.reset", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyDetailsConfigurationForcecommonscopeforall(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.company.details.configuration.forcecommonscopeforall", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
