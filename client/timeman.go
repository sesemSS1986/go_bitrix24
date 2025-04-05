package client

func (c *Client) TimemanSettings(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "timeman.settings", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanStatus(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "timeman.status", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanOpen(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "timeman.open", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanClose(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "timeman.close", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanPause(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "timeman.pause", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
