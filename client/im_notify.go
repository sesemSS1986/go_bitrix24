package client

func (c *Client) ImNotifyPersonalAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.notify.personal.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImNotifySystemAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.notify.system.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImNotifyDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.notify.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImNotifyRead(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.notify.read", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
