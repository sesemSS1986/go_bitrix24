package client

func (c *Client) ImbotAppRegister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.app.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotAppUnregister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.app.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotAppUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.app.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
