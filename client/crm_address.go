package client

func (c *Client) CrmAddressFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.address.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAddressAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.address.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAddressUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.address.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAddressList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.address.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAddressDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.address.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
