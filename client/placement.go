package client

func (c *Client) PlacementList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"placement.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) PlacementBind(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"placement.bind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) PlacementUnbind(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"placement.unbind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) PlacementGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"placement.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
