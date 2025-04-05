package client

func (c *Client) ImbotCommandRegister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.command.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotCommandUnregister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.command.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotCommandUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.command.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotCommandAnswer(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imbot.command.answer", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
