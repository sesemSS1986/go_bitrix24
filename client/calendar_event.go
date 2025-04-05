package client

func (c *Client) CalendarEventGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "calendar.event.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarEventAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "calendar.event.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarEventUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "calendar.event.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarEventDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "calendar.event.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarEventGetNearest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "calendar.event.get.nearest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarEventGetbyid(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "calendar.event.getbyid", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
