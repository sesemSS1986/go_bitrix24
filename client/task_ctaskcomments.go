package client

func (c *Client) TaskCtaskcommentsGetmanifest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskcomments.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskcommentsAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctaskcomments.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
