package client

func (c *Client) CrmQuoteFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteProductrowsSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.productrows.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteProductrowsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.productrows.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteContactFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.contact.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteContactAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.contact.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteContactDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.contact.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteContactItemsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.contact.items.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteContactItemsSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.contact.items.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteContactItemsDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.contact.items.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteUserfieldAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteUserfieldGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteUserfieldList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.userfield.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteUserfieldUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteUserfieldDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.quote.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
