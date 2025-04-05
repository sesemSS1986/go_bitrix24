package client

func (c *Client) SmileGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "smile.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
