package client

func (c *Client) DiskVersionGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.version.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
