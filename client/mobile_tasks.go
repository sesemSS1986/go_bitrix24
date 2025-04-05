package client

func (c *Client) MobileTasksDeadlinesGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.tasks.deadlines.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
