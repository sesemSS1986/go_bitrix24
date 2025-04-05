package client

func (c *Client) CalendarEventGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.event.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarEventAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.event.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarEventUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.event.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarEventDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.event.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarEventGetNearest(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.event.get.nearest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarEventGetbyid(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.event.getbyid", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
