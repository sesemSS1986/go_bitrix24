package client

func (c *Client) EntityAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityRights(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.rights", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
