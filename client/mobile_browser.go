package client

func (c *Client) MobileBrowserConstGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.browser.const.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
