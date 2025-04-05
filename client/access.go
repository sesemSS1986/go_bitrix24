package client

func (c *Client) AccessName(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "access.name", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
