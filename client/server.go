package client

func (c *Client) ServerTime(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("server.time", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
