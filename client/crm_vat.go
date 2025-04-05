package client

func (c *Client) CrmVatFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.vat.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmVatList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.vat.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmVatGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.vat.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmVatAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.vat.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmVatUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.vat.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmVatDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.vat.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
