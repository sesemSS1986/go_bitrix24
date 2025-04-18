package client

func (c *Client) ImMobileConfigGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.mobile.config.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
