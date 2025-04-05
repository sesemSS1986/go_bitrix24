package client

func (c *Client) ImbotAppRegister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imbot.app.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotAppUnregister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imbot.app.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotAppUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imbot.app.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
