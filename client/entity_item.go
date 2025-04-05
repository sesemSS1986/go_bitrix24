package client

func (c *Client) EntityItemAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.item.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityItemGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.item.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityItemUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.item.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityItemDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.item.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityItemPropertyAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.item.property.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityItemPropertyGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.item.property.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityItemPropertyUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.item.property.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityItemPropertyDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.item.property.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
