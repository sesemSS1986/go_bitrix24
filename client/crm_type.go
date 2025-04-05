package client

func (c *Client) CrmTypeFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.type.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTypeAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.type.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTypeGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.type.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTypeList(root bool, url string, p Parameters) (map[string]interface{}, error) {
	resp, err := c.Request(root, url, "crm.type.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTypeUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.type.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmTypeDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.type.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
