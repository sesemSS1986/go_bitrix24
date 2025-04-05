package client

func (c *Client) CalendarSectionGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.section.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarSectionAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.section.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarSectionUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.section.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarSectionDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.section.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
