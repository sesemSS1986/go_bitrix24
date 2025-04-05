package client

func (c *Client) CrmItemFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.item.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmItemAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.item.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmItemGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.item.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmItemList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.item.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmItemUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.item.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmItemDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.item.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmItemImport(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.item.import", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmItemBatchImport(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.item.batchImport", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
