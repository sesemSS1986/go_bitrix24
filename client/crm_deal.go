package client

func (c *Client) CrmDealFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealProductrowsSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.productrows.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealProductrowsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.productrows.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealContactFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.contact.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealContactAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.contact.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealContactDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.contact.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealContactItemsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.contact.items.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealContactItemsSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.contact.items.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealContactItemsDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.contact.items.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealRecurringFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.recurring.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealRecurringList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.recurring.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealRecurringAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.recurring.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealRecurringGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.recurring.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealRecurringUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.recurring.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealRecurringDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.recurring.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealRecurringExpose(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.recurring.expose", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealUserfieldAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealUserfieldGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealUserfieldList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.userfield.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealUserfieldUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealUserfieldDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealDetailsConfigurationGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.details.configuration.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealDetailsConfigurationSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.details.configuration.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealDetailsConfigurationReset(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.details.configuration.reset", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmDealDetailsConfigurationForcecommonscopeforall(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.deal.details.configuration.forcecommonscopeforall", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
