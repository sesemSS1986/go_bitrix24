package client

func (c *Client) Sonet_groupUserGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupUserInvite(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group.user.invite", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupUserRequest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group.user.request", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupUserAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group.user.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupUserUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group.user.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupUserDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group.user.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupUserGroups(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group.user.groups", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
