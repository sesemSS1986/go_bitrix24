package client

func (c *Client) TimemanSettings(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.settings", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanStatus(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.status", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanOpen(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.open", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanClose(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.close", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanPause(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.pause", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
