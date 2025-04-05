package client

import "github.com/sesemSS1986/go_bitrix24/types"

func (c *Client) TimemanTimecontrolSettingsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.timecontrol.settings.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanTimecontrolSettingsSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.timecontrol.settings.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanTimecontrolReportAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.timecontrol.report.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanTimecontrolReportsSettingsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.timecontrol.reports.settings.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanTimecontrolReportsUsersGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.timecontrol.reports.users.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanTimecontrolReportsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.timecontrol.reports.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanTimecontrolReport(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("timeman.timecontrol.report", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
