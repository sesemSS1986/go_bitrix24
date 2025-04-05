package client

func (c *Client) BizprocActivityAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"bizproc.activity.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocActivityUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"bizproc.activity.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocActivityDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"bizproc.activity.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocActivityLog(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"bizproc.activity.log", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocActivityList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"bizproc.activity.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
