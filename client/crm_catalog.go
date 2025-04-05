package client

func (c *Client) CrmCatalogFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.catalog.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCatalogGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.catalog.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmCatalogList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.catalog.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
