package client

func (c *Client) ImUserGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserListGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.user.list.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserBusinessList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.user.business.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserBusinessGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.user.business.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserStatusGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.user.status.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserStatusSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.user.status.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserStatusIdleStart(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.user.status.idle.start", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserStatusIdleEnd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.user.status.idle.end", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImUserStatusIdleContinue(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.user.status.idle.continue", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
