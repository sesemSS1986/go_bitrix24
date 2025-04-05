package client

func (c *Client) _placements(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"_placements", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
