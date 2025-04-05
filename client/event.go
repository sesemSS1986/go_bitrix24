package client

func (c *Client) EventBind(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "event.bind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EventUnbind(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "event.unbind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EventGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "event.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EventOfflineGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "event.offline.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EventOfflineClear(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "event.offline.clear", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EventOfflineError(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "event.offline.error", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EventOfflineList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "event.offline.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EventTest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "event.test", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
