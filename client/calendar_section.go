package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) CalendarSectionGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.section.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarSectionAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.section.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarSectionUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.section.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CalendarSectionDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("calendar.section.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
