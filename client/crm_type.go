package client

func (c *Client) CrmTypeFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.type.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTypeAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.type.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTypeGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.type.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTypeList(p Parameters) (map[string]interface{}, error) {
	resp, err := c.Request("crm.type.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTypeUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.type.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTypeDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.type.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
