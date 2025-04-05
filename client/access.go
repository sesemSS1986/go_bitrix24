package client

func (c *Client) AccessName(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("access.name", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
