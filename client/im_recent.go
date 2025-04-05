package client

func (c *Client) ImRecentGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.recent.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImRecentList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.recent.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImRecentPin(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.recent.pin", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImRecentHide(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.recent.hide", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImRecentUnread(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.recent.unread", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
