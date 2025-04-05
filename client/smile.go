package client

func (c *Client) SmileGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("smile.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
