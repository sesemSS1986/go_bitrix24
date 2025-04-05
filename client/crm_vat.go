package client

func (c *Client) CrmVatFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.vat.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmVatList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.vat.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmVatGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.vat.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmVatAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.vat.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmVatUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.vat.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmVatDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.vat.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
