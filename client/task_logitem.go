package client

func (c *Client) TaskLogitemGetmanifest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.logitem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskLogitemList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.logitem.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
