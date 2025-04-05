package client

func (c *Client) CrmProductFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyTypes(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.property.types", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.property.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertySettingsFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.property.settings.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyEnumerationFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.property.enumeration.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.property.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.property.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.property.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.property.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.product.property.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
