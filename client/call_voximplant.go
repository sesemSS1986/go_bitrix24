package client

func (c *Client) VoximplantCallbackStart(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "voximplant.callback.start", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantInfocallStartwithtext(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "voximplant.infocall.startwithtext", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) VoximplantInfocallStartwithsound(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "voximplant.infocall.startwithsound", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
