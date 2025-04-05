package client

func (c *Client) MobileSettingsTabsSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "mobile.settings.tabs.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileSettingsEnergySet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "mobile.settings.energy.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileSettingsEnergyGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "mobile.settings.energy.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileSettingsOtherSet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "mobile.settings.other.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileSettingsOtherGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "mobile.settings.other.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
