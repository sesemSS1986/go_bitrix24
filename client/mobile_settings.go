package client

func (c *Client) MobileSettingsTabsSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.settings.tabs.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileSettingsEnergySet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.settings.energy.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileSettingsEnergyGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.settings.energy.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileSettingsOtherSet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.settings.other.set", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MobileSettingsOtherGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mobile.settings.other.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
