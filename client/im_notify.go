package client

func (c *Client) ImNotifyPersonalAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.notify.personal.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImNotifySystemAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.notify.system.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImNotifyDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.notify.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImNotifyRead(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.notify.read", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
