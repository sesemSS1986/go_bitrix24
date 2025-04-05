package client

func (c *Client) CalendarUserSettingsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.user.settings.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarUserSettingsSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.user.settings.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
