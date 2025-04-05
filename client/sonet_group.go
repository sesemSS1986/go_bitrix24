package client

func (c *Client) Sonet_groupGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sonet_group.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupCreate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sonet_group.create", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sonet_group.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sonet_group.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_groupSetowner(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "sonet_group.setowner", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
