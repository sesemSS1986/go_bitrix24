package client

func (c *Client) ImMobileConfigGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.mobile.config.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
