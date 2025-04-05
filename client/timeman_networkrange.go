package client

func (c *Client) TimemanNetworkrangeGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.networkrange.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanNetworkrangeSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.networkrange.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanNetworkrangeCheck(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.networkrange.check", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
