package client

func (c *Client) Sonet_group_subjectGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"sonet_group_subject.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_group_subjectAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"sonet_group_subject.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_group_subjectUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"sonet_group_subject.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Sonet_group_subjectDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"sonet_group_subject.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
