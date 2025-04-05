package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CalendarSettingsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.settings.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
