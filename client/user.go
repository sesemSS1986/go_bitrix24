package client

func (c *Client) UserOptionGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("user.option.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserOptionSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("user.option.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserAdmin(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("user.admin", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserAccess(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("user.access", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("user.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserCurrent(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("user.current", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserSearch(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("user.search", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("user.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("user.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserOnline(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("user.online", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserCounters(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("user.counters", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
