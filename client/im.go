package client

func (c *Client) ImNotify(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.notify", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
