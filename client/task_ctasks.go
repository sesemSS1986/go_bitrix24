package client

func (c *Client) TaskCtasksGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.ctasks.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
