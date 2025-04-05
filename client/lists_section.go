package client

func (c *Client) ListsSectionAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"lists.section.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsSectionGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"lists.section.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsSectionUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"lists.section.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsSectionDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"lists.section.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
