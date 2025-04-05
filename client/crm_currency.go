package client

func (c *Client) CrmCurrencyFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.currency.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.currency.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.currency.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.currency.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.currency.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.currency.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyLocalizationsFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.currency.localizations.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyLocalizationsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.currency.localizations.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyLocalizationsSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.currency.localizations.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyLocalizationsDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.currency.localizations.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyBaseSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.currency.base.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCurrencyBaseGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.currency.base.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
