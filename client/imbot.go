package client

func (c *Client) ImbotRegister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotUnregister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotSendtyping(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.sendtyping", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
