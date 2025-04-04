package client

func (c *Client) MobileBrowserConstGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.browser.const.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
