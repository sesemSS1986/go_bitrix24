package client

func (c *Client) Batch(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"batch", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Scope(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"scope", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Events(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"events", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Profile(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"profile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
