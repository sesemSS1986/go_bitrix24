package client

func (c *Client) MobilePushTypesGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.push.types.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobilePushConfigGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.push.config.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobilePushConfigSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.push.config.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobilePushStatusGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.push.status.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobilePushStatusSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.push.status.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobilePushSmartfilterStatusGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.push.smartfilter.status.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobilePushSmartfilterStatusSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.push.smartfilter.status.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
