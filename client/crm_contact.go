package client

func (c *Client) CrmContactFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactCompanyFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.company.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactCompanyAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.company.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactCompanyDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.company.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactCompanyItemsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.company.items.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactCompanyItemsSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.company.items.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactCompanyItemsDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.company.items.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactUserfieldAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactUserfieldGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactUserfieldList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.userfield.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactUserfieldUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactUserfieldDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactDetailsConfigurationGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.details.configuration.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactDetailsConfigurationSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.details.configuration.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactDetailsConfigurationReset(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.details.configuration.reset", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactDetailsConfigurationForcecommonscopeforall(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.contact.details.configuration.forcecommonscopeforall", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
