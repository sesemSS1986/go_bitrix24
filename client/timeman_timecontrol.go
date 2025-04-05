package client

func (c *Client) TimemanTimecontrolSettingsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"timeman.timecontrol.settings.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanTimecontrolSettingsSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"timeman.timecontrol.settings.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanTimecontrolReportAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"timeman.timecontrol.report.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanTimecontrolReportsSettingsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"timeman.timecontrol.reports.settings.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanTimecontrolReportsUsersGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"timeman.timecontrol.reports.users.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanTimecontrolReportsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"timeman.timecontrol.reports.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TimemanTimecontrolReport(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"timeman.timecontrol.report", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
