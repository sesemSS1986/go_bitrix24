package client

func (c *Client) ImNotify(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.notify", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
