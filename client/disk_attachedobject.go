package client

func (c *Client) DiskAttachedobjectGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.attachedobject.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
