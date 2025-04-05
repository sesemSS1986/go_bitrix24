package client

func (c *Client) TelephonyExternalcallSearchcrmentities(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("telephony.externalcall.searchcrmentities", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternalcallRegister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("telephony.externalcall.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternalcallFinish(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("telephony.externalcall.finish", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternalcallShow(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("telephony.externalcall.show", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternalcallHide(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("telephony.externalcall.hide", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TelephonyExternalcallAttachrecord(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("telephony.externalcall.attachrecord", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
