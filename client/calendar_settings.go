package client

func (c *Client) CalendarSettingsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.settings.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
