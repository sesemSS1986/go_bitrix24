package client

func (c *Client) ListsAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"lists.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"lists.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"lists.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"lists.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
