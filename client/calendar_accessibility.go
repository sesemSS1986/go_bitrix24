package client

func (c *Client) CalendarAccessibilityGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.accessibility.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
