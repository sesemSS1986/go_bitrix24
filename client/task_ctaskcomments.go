package client

func (c *Client) TaskCtaskcommentsGetmanifest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.ctaskcomments.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskcommentsAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.ctaskcomments.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
