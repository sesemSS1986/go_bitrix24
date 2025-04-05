package client

func (c *Client) PullApplicationConfigGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"pull.application.config.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) PullApplicationEventAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"pull.application.event.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) PullApplicationPushAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"pull.application.push.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
