package client

func (c *Client) TaskCtasksGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.ctasks.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
