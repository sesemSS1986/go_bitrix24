package client

func (c *Client) CrmCompanyFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyContactFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.contact.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyContactAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.contact.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyContactDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.contact.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyContactItemsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.contact.items.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyContactItemsSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.contact.items.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyContactItemsDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.contact.items.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyUserfieldAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyUserfieldGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyUserfieldList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.userfield.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyUserfieldUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyUserfieldDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyDetailsConfigurationGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.details.configuration.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyDetailsConfigurationSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.details.configuration.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyDetailsConfigurationReset(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.details.configuration.reset", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCompanyDetailsConfigurationForcecommonscopeforall(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.company.details.configuration.forcecommonscopeforall", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
