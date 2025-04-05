package client

func (c *Client) ListsAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
