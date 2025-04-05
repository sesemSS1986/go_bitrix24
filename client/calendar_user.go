package client

func (c *Client) CalendarUserSettingsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.user.settings.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarUserSettingsSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.user.settings.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
