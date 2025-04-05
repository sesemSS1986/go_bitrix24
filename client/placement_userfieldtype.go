package client

func (c *Client) UserfieldtypeList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("userfieldtype.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserfieldtypeAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("userfieldtype.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserfieldtypeUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("userfieldtype.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserfieldtypeDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("userfieldtype.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
