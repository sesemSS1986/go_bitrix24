package client

func (c *Client) CrmProductsectionFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.productsection.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductsectionAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.productsection.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductsectionGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.productsection.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductsectionList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.productsection.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductsectionUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.productsection.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductsectionDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.productsection.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
