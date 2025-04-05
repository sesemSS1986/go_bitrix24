package client

func (c *Client) ImRecentGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.recent.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImRecentList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.recent.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImRecentPin(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.recent.pin", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImRecentHide(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.recent.hide", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImRecentUnread(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.recent.unread", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
