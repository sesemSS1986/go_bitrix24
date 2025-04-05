package client

func (c *Client) BizprocTaskList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"bizproc.task.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocTaskComplete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"bizproc.task.complete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
