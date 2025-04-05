package client

func (c *Client) CrmProductrowFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.productrow.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductrowAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.productrow.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductrowGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.productrow.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductrowList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.productrow.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductrowUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.productrow.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductrowDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.productrow.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
