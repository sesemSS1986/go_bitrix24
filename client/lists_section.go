package client

func (c *Client) ListsSectionAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.section.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsSectionGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.section.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsSectionUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.section.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsSectionDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.section.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
