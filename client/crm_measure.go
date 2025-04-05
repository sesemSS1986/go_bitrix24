package client

func (c *Client) CrmMeasureFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.measure.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmMeasureAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.measure.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmMeasureGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.measure.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmMeasureList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.measure.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmMeasureUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.measure.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmMeasureDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.measure.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
