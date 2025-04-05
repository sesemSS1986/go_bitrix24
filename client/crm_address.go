package client

func (c *Client) CrmAddressFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.address.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAddressAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.address.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAddressUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.address.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAddressList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.address.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmAddressDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.address.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
