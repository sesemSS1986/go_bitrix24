package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CalendarMeetingStatusSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.meeting.status.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarMeetingParamsSet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.meeting.params.set", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}

func (c *Client) CalendarMeetingStatusGet(data interface{}) (*types.Response, error) {
	resp, err := c.DoRaw("calendar.meeting.status.get", data, &types.Response{})
	if err != nil {
		return nil, err
	}
	return resp.Result().(*types.Response), err
}
