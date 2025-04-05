package client

func (c *Client) MobileFormProfile(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "mobile.form.profile", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
