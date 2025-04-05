package client

func (c *Client) BizprocActivityAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.activity.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocActivityUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.activity.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocActivityDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.activity.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocActivityLog(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.activity.log", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) BizprocActivityList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("bizproc.activity.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
