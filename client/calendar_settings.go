package client

func (c *Client) CalendarSettingsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "calendar.settings.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
