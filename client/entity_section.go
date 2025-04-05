package client

func (c *Client) EntitySectionAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.section.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntitySectionGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.section.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntitySectionUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.section.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EntitySectionDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("entity.section.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
