package client

func (c *Client) MobileUserGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "mobile.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileUserCanusetelephony(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "mobile.user.canusetelephony", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
