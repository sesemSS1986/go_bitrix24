package client

func (c *Client) ServerTime(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"server.time", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
