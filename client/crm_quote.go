package client

func (c *Client) CrmQuoteFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteProductrowsSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.productrows.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteProductrowsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.productrows.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteContactFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.contact.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteContactAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.contact.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteContactDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.contact.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteContactItemsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.contact.items.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteContactItemsSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.contact.items.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteContactItemsDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.contact.items.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteUserfieldAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.userfield.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteUserfieldGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.userfield.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteUserfieldList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.userfield.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteUserfieldUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.userfield.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmQuoteUserfieldDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.quote.userfield.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
