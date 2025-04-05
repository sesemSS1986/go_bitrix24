package client

func (c *Client) CalendarResourceList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.resource.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarResourceAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.resource.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarResourceUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.resource.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarResourceDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.resource.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarResourceBookingList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.resource.booking.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
