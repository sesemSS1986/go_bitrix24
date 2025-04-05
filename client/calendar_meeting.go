package client

func (c *Client) CalendarMeetingStatusSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.meeting.status.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarMeetingParamsSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.meeting.params.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarMeetingStatusGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.meeting.status.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
