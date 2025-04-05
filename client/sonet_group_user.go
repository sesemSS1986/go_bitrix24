package client

func (c *Client) Sonet_groupUserGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"sonet_group.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupUserInvite(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"sonet_group.user.invite", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupUserRequest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"sonet_group.user.request", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupUserAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"sonet_group.user.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupUserUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"sonet_group.user.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupUserDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"sonet_group.user.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupUserGroups(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"sonet_group.user.groups", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
