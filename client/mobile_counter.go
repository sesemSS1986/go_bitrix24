package client

func (c *Client) MobileCounterTypesGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.counter.types.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileCounterGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.counter.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileCounterConfigGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.counter.config.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileCounterConfigSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.counter.config.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
