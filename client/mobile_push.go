package client

func (c *Client) MobilePushTypesGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.push.types.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobilePushConfigGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.push.config.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobilePushConfigSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.push.config.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobilePushStatusGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.push.status.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobilePushStatusSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.push.status.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobilePushSmartfilterStatusGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.push.smartfilter.status.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobilePushSmartfilterStatusSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.push.smartfilter.status.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
