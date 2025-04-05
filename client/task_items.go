package client

func (c *Client) TaskItemsGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "task.items.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
