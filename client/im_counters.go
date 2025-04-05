package client

func (c *Client) ImCountersGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.counters.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
