package client

func (c *Client) Tasks_extendedMetaSetanystatus(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"tasks_extended.meta.setanystatus", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Tasks_extendedMetaOccurinlogsas(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"tasks_extended.meta.occurinlogsas", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
