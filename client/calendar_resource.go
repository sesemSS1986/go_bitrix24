package client

func (c *Client) CalendarResourceList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.resource.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarResourceAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.resource.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarResourceUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.resource.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarResourceDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.resource.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarResourceBookingList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.resource.booking.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
