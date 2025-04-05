package client

func (c *Client) _placements(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("_placements", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
