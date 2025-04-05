package client

func (c *Client) CrmProductsectionFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.productsection.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductsectionAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.productsection.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductsectionGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.productsection.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductsectionList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.productsection.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductsectionUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.productsection.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductsectionDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.productsection.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
