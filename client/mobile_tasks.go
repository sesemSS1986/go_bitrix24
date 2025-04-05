package client

func (c *Client) MobileTasksDeadlinesGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mobile.tasks.deadlines.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
