package client

func (c *Client) AppOptionGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "app.option.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) AppOptionSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "app.option.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) AppInfo(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "app.info", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
