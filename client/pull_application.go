package client

func (c *Client) PullApplicationConfigGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("pull.application.config.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) PullApplicationEventAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("pull.application.event.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) PullApplicationPushAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("pull.application.push.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
