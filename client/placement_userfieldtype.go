package client

func (c *Client) UserfieldtypeList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"userfieldtype.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserfieldtypeAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"userfieldtype.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserfieldtypeUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"userfieldtype.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserfieldtypeDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"userfieldtype.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
