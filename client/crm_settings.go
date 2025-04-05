package client

func (c *Client) CrmSettingsModeGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.settings.mode.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
