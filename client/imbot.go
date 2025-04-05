package client

func (c *Client) ImbotRegister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotUnregister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.unregister", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotSendtyping(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imbot.sendtyping", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
