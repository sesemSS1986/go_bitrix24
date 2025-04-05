package client

func (c *Client) TelephonyExternallineAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"telephony.externalline.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternallineUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"telephony.externalline.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternallineDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"telephony.externalline.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternallineGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"telephony.externalline.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
