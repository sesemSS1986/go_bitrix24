package client

func (c *Client) AppOptionGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("app.option.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) AppOptionSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("app.option.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) AppInfo(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("app.info", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
