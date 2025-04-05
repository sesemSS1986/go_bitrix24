package client

func (c *Client) TelephonyExternalcallSearchcrmentities(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"telephony.externalcall.searchcrmentities", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternalcallRegister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"telephony.externalcall.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternalcallFinish(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"telephony.externalcall.finish", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternalcallShow(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"telephony.externalcall.show", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternalcallHide(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"telephony.externalcall.hide", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternalcallAttachrecord(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"telephony.externalcall.attachrecord", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
