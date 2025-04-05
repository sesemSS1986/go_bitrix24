package client

func (c *Client) CalendarAccessibilityGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.accessibility.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
