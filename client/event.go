package client

func (c *Client) EventBind(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("event.bind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EventUnbind(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("event.unbind", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EventGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("event.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EventOfflineGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("event.offline.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EventOfflineClear(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("event.offline.clear", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EventOfflineError(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("event.offline.error", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EventOfflineList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("event.offline.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) EventTest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("event.test", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
