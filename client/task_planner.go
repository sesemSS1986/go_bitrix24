package client

func (c *Client) TaskPlannerGetcurrenttaskslist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.planner.getcurrenttaskslist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskPlannerGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.planner.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
