package client

func (c *Client) TimemanNetworkrangeGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "timeman.networkrange.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanNetworkrangeSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "timeman.networkrange.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanNetworkrangeCheck(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "timeman.networkrange.check", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
