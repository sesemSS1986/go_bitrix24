package client

func (c *Client) MobileCounterTypesGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.counter.types.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileCounterGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.counter.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileCounterConfigGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.counter.config.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileCounterConfigSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.counter.config.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
