package client

func (c *Client) DiskRightsGettasks(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("disk.rights.gettasks", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
