package client

func (c *Client) EntitySectionAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.section.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntitySectionGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.section.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntitySectionUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.section.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntitySectionDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "entity.section.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
