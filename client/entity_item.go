package client

func (c *Client) EntityItemAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.item.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityItemGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.item.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityItemUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.item.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityItemDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.item.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityItemPropertyAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.item.property.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityItemPropertyGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.item.property.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityItemPropertyUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.item.property.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityItemPropertyDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.item.property.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
