package client

func (c *Client) UserOptionGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"user.option.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserOptionSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"user.option.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserAdmin(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"user.admin", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserAccess(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"user.access", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"user.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserCurrent(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"user.current", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserSearch(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"user.search", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"user.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"user.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserOnline(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"user.online", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserCounters(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"user.counters", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
