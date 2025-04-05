package client

func (c *Client) EntityAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntityRights(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.rights", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
