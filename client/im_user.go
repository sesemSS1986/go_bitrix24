package client

func (c *Client) ImUserGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserListGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.user.list.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserBusinessList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.user.business.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserBusinessGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.user.business.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserStatusGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.user.status.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserStatusSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.user.status.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserStatusIdleStart(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.user.status.idle.start", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserStatusIdleEnd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.user.status.idle.end", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserStatusIdleContinue(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.user.status.idle.continue", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
