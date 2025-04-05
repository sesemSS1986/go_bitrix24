package client

func (c *Client) CrmProductFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyTypes(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.property.types", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.property.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertySettingsFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.property.settings.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyEnumerationFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.property.enumeration.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.property.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.property.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.property.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.property.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmProductPropertyDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.product.property.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
