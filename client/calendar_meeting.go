package client

func (c *Client) CalendarMeetingStatusSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.meeting.status.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarMeetingParamsSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.meeting.params.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarMeetingStatusGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"calendar.meeting.status.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
