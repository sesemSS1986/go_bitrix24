package client

func (c *Client) BizprocProviderAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"bizproc.provider.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocProviderDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"bizproc.provider.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocProviderList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"bizproc.provider.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
