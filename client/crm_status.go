package client

func (c *Client) CrmStatusFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.status.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.status.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.status.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.status.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.status.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.status.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusEntityTypes(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.status.entity.types", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusEntityItems(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.status.entity.items", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmStatusExtraFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.status.extra.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
