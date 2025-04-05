package client

func (c *Client) TaskItemsGetlist(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("task.items.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
