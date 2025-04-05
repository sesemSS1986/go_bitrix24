package client

func (c *Client) CrmContactFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactCompanyFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.company.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactCompanyAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.company.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactCompanyDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.company.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactCompanyItemsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.company.items.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactCompanyItemsSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.company.items.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactCompanyItemsDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.company.items.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactUserfieldAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactUserfieldGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactUserfieldList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.userfield.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactUserfieldUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactUserfieldDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactDetailsConfigurationGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.details.configuration.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactDetailsConfigurationSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.details.configuration.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactDetailsConfigurationReset(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.details.configuration.reset", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmContactDetailsConfigurationForcecommonscopeforall(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.contact.details.configuration.forcecommonscopeforall", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
