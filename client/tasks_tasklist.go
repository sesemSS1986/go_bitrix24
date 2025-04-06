package client

func (c *Client) TaskTaskList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "tasks.task.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
